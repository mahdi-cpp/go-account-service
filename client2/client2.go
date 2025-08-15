package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6389", // Connect to your custom RESP server
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to RESP server: %v", err)
	}
	fmt.Println("Connected to custom RESP server using go-redis!")

	ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Add a timeout for publish
	//_, err = rdb.Publish(ctx2, "channel_2", "send by client2").Result()
	_, err = rdb.Publish(ctx2, "account/command", "list").Result()
	cancel() // Release resources associated with the context

	if err != nil {
		log.Printf("[Publisher] Error publishing: %v", err)
	}

	fmt.Println("Main application waiting for goroutines...")
	time.Sleep(5 * time.Second) // Keep main alive to see subscriber output

	err = rdb.Close()
	if err != nil {
		return
	}

	fmt.Println("Main application exiting.")
}
