package account

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mahdi-cpp/api-go-pkg/collection_manager_uuid7"
	"github.com/mahdi-cpp/go-account-service/utils"
	"github.com/redis/go-redis/v9"
)

const (
	commandChannel      = "account/command"
	userChannel         = "account/user"
	listChannel         = "account/list"
	userAddChannel      = "account/user/add"
	userDeleteChannel   = "account/user/delete"
	userUpdateChannel   = "account/user/update"
	publishTimeout      = 2 * time.Second
	subscriptionTimeout = 3 * time.Second
)

type ServiceManager struct {
	mu             sync.RWMutex
	UserCollection *collection_manager_uuid7.Manager[*User]
	rdb            *redis.Client
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
	subReady       chan struct{}
}

func NewAccountManager() (*ServiceManager, error) {

	ctx, cancel := context.WithCancel(context.Background())
	manager := &ServiceManager{
		ctx:      ctx,
		cancel:   cancel,
		subReady: make(chan struct{}),
	}

	// Initialize Redis client
	manager.rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6389",
		DB:   0,
	})

	// Verify Redis connection
	ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if _, err := manager.rdb.Ping(ctxPing).Result(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	// Initialize user collection
	var err error
	manager.UserCollection, err = collection_manager_uuid7.NewCollectionManager[*User](
		GetServicesPath("account/users/"), true,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user collection: %w", err)
	}

	// Start subscription handler
	manager.wg.Add(1)
	go manager.runSubscription()

	// Wait for subscription to be ready
	select {
	case <-manager.subReady:
		log.Println("Redis subscription established")
	case <-time.After(subscriptionTimeout):
		log.Println("Warning: Subscription setup timed out")
	case <-ctx.Done():
		return nil, context.Canceled
	}

	return manager, nil
}

func (m *ServiceManager) Close() error {
	m.cancel()  // Signal shutdown
	m.wg.Wait() // Wait for goroutines

	if err := m.rdb.Close(); err != nil {
		return fmt.Errorf("redis close error: %w", err)
	}
	return nil
}

func (m *ServiceManager) runSubscription() {
	defer m.wg.Done()

	channels := []string{
		commandChannel,
		userChannel,
		listChannel,
		userAddChannel,
		userDeleteChannel,
		userUpdateChannel,
	}

	pubsub := m.rdb.Subscribe(m.ctx, channels...)
	defer pubsub.Close()

	// Confirm subscription
	if _, err := pubsub.ReceiveTimeout(m.ctx, 500*time.Millisecond); err != nil {
		log.Printf("Subscription confirmation failed: %v", err)
		return
	}
	close(m.subReady)

	ch := pubsub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.Println("Subscription channel closed")
				return
			}
			m.handleMessage(msg)
		case <-m.ctx.Done():
			log.Println("Subscription exiting due to shutdown")
			return
		}
	}
}

func (m *ServiceManager) handleMessage(msg *redis.Message) {
	switch msg.Channel {
	case commandChannel:
		switch msg.Payload {
		case "list":
			if err := m.publish(); err != nil {
				log.Printf("Publish failed: %v", err)
			}
		case "user":
			// Handle user command
		}
	case userAddChannel:
		// Handle user addition
	case userDeleteChannel:
		// Handle user deletion
	case userUpdateChannel:
		// Handle user update
	}
}

func (m *ServiceManager) publish() error {
	users, err := m.UserCollection.GetAll()
	if err != nil {
		return fmt.Errorf("get users failed: %w", err)
	}

	toJSON, err := utils.ToStringJson(users)
	if err != nil {
		return fmt.Errorf("JSON conversion failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(m.ctx, publishTimeout)
	defer cancel()

	if err := m.rdb.Publish(ctx, listChannel, toJSON).Err(); err != nil {
		return fmt.Errorf("redis publish failed: %w", err)
	}

	return nil
}
