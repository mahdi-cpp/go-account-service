package main

import (
	"fmt"
	"github.com/mahdi-cpp/go-account-service/account"
	"log"
)

func main() {

	ginInit()

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

	manager.Publish()

	h := account.NewUserHandler(manager)
	userRoute(h)

	startServer(router)

}

func userRoute(h *account.UserHandler) {

	api := router.Group("/api/v1/user")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("get_user", h.GetUser)
	api.POST("list", h.GetList)
}
