package main

import (
	"fmt"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {

	// Create manager
	manager, err := account.NewClientManager()
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Close() // Ensure cleanup

	// Register callback
	manager.Register(func(msg *redis.Message) {

		fmt.Println("AccountService:", msg.Channel)

		switch msg.Channel {
		case "user":

		}
		for _, user := range manager.Users {
			//fmt.Println(user.ID)
			fmt.Println(user.FirstName, user.LastName, "   ", user.PhoneNumber)
			fmt.Println("---------------------------")
		}
	})

	// Request data
	if err := manager.RequestList(); err != nil {
		log.Printf("Error requesting list: %v", err)
	}

	// Start additional subscribers
	manager.StartSubscriber("account/notifications", "account/alerts")

	fmt.Println("12")

	select {}
}
