package depricated

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var ctx = context.Background()

type ClientManager struct {
	mu       sync.RWMutex
	rdb      *redis.Client
	Users    []*account.User
	callback func(msg *redis.Message)
}

func (manager *ClientManager) Register(callback func(msg *redis.Message)) {
	manager.callback = callback
}

func NewClientManager() *ClientManager {
	manager := &ClientManager{}

	manager.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6389", // Connect to your RESP Broker
		Password: "",
		DB:       0,
	})
	//defer manager.rdb.Close()

	_, err := manager.rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to RESP server: %v", err)
	}
	fmt.Println("Connected to custom RESP server using go-redis!")

	go func() {

		pubsub := manager.rdb.Subscribe(ctx,
			"account/list",
		)
		defer pubsub.Close()

		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println("Subscribe")
			if msg.Channel == "account/user" {
				//manager.userChanel(msg.Payload)
			} else if msg.Channel == "account/list" {
				manager.fetchUsers(msg)
			}
		}

		fmt.Println("[Subscriber] Goroutine finished.")
	}()
	return manager
}

func (manager *ClientManager) RequestList() error {

	time.Sleep(20 * time.Millisecond)

	ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Add a timeout for publish
	_, err := manager.rdb.Publish(ctx2, "account/command", "list").Result()
	cancel() // Release resources associated with the context

	if err != nil {
		log.Printf("[Publisher] Error publishing: %v", err)
		return err
	}
	log.Println("Request List ok")

	return nil
}

func (manager *ClientManager) fetchUsers(msg *redis.Message) {

	fmt.Println("fetchUsers")
	err := json.Unmarshal([]byte(msg.Payload), &manager.Users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	manager.callback(msg)
}

func (manager *ClientManager) StartAnotherSubscriber(channels ...string) {
	go func() {
		pubsub := manager.rdb.Subscribe(ctx, channels...)
		defer pubsub.Close()
		ch := pubsub.Channel()

		fmt.Printf("[Subscriber] Listening on channels: %v\n", channels)
		for msg := range ch {
			fmt.Printf("[Subscriber %v] Received message from %s: %s\n", channels, msg.Channel, msg.Payload)
		}
		fmt.Printf("[Subscriber] Goroutine for %v finished.\n", channels)
	}()
}
