package main

import (
	"fmt"
	"github.com/mahdi-cpp/go-account-service/account"
	"log"
)

func main() {

	// Create account manager
	manager, err := account.NewAccountManager()
	if err != nil {
		log.Fatalf("Failed to create manager: %v", err)
	}
	defer manager.Close() // Ensure proper cleanup

	users, err := manager.UserCollection.GetAll()
	if err != nil {
		return
	}

	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName)
	}

	fmt.Println("start listening")

	select {}
}
