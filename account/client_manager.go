package account

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type ClientManager struct {
	mu        sync.RWMutex
	rdb       *redis.Client
	Users     []*User
	UsersMap  map[string]*User
	callback  func(msg *redis.Message)
	ctx       context.Context
	cancel    context.CancelFunc
	closeOnce sync.Once
	subReady  chan struct{} // Signals when main subscription is ready
}

func NewClientManager() (*ClientManager, error) {
	ctx, cancel := context.WithCancel(context.Background())
	manager := &ClientManager{
		rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:6389",
			DB:   0,
		}),
		UsersMap: make(map[string]*User),
		ctx:      ctx,
		cancel:   cancel,
		subReady: make(chan struct{}),
	}

	// Verify Redis connection with timeout
	ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := manager.rdb.Ping(ctxPing).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	go manager.runMainSubscription()
	return manager, nil
}

func (m *ClientManager) GetUsersMap() map[string]*User {
	return m.UsersMap
}

func (m *ClientManager) runMainSubscription() {
	pubsub := m.rdb.Subscribe(m.ctx, listChannel)
	defer pubsub.Close()

	// Wait for subscription confirmation
	if _, err := pubsub.Receive(m.ctx); err != nil {
		log.Printf("Failed to receive subscription confirmation: %v", err)
		return
	}

	// Signal that subscription is ready
	close(m.subReady)

	ch := pubsub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.Printf("Main subscription channel closed")
				return
			}
			m.handleMessage(msg)
		case <-m.ctx.Done():
			log.Printf("Main subscription exiting due to context cancellation")
			return
		}
	}
}

func (m *ClientManager) Close() {
	m.closeOnce.Do(func() {
		m.cancel() // Signal shutdown to all goroutines
		if err := m.rdb.Close(); err != nil {
			log.Printf("Redis client close error: %v", err)
		}
	})
}

func (m *ClientManager) Register(callback func(msg *redis.Message)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.callback = callback
}

func (m *ClientManager) RequestList() error {
	// Wait for subscription to be ready or timeout
	select {
	case <-m.subReady:
		// Subscription is ready, proceed
	case <-time.After(100 * time.Millisecond):
		log.Println("Warning: Subscription not ready after timeout")
	case <-m.ctx.Done():
		return context.Canceled
	}

	ctx, cancel := context.WithTimeout(m.ctx, 2*time.Second)
	defer cancel()

	if err := m.rdb.Publish(ctx, commandChannel, "list").Err(); err != nil {
		return fmt.Errorf("publish failed: %w", err)
	}
	return nil
}

func (m *ClientManager) StartSubscriber(channels ...string) {
	go m.runSubscription(channels)
}

func (m *ClientManager) runSubscription(channels []string) {
	pubsub := m.rdb.Subscribe(m.ctx, channels...)
	defer pubsub.Close()

	// Wait for subscription confirmation
	if _, err := pubsub.Receive(m.ctx); err != nil {
		log.Printf("Failed to subscribe to %v: %v", channels, err)
		return
	}

	ch := pubsub.Channel()
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				log.Printf("Subscription closed: %v", channels)
				return
			}
			m.handleMessage(msg)
		case <-m.ctx.Done():
			log.Printf("Shutting down subscriber: %v", channels)
			return
		}
	}
}

func (m *ClientManager) handleMessage(msg *redis.Message) {
	switch msg.Channel {
	case listChannel:
		m.fetchUsers(msg)
	default:
		log.Printf("Received message on %s: %s", msg.Channel, msg.Payload)
	}
}

func (m *ClientManager) fetchUsers(msg *redis.Message) {
	var users []*User
	if err := json.Unmarshal([]byte(msg.Payload), &users); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return
	}

	m.mu.Lock()
	m.Users = users
	m.mu.Unlock()

	m.mu.RLock()
	cb := m.callback
	m.mu.RUnlock()

	if cb != nil {
		cb(msg)
	}
}
