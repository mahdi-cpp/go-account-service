package main

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-pkg/collection_manager_uuid7"
	"github.com/mahdi-cpp/api-go-pkg/metadata"
	"github.com/mahdi-cpp/api-go-pkg/network"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/mahdi-cpp/go-account-service/depricated"
	"log"
)

func testAccount(id int) {

}

func testAccountUserList() {

	ac := depricated.NewNetworkAccountManager()
	users, err := ac.GetAll()
	if err != nil {
		return
	}

	for _, user := range *users {
		fmt.Printf("User ID: %d\n", user.ID)
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
	}

	//var name = asset.PHAsset{}

	//name.Update()

}

func testCollection2() {
	users, err := collection_manager_uuid7.NewCollectionManager[*account.User]("albums_test.json", false)
	if err != nil {
		fmt.Println("UserStorage:", err)
		return
	}

	item := &account.User{FirstName: "Original"}
	create, err := users.Create(item)
	if err != nil {
		return
	}

	fmt.Println(create.FirstName)

	//fmt.Println(createdItem.ID))
}

func testInfoPlist() {
	infoPlist := metadata.NewMetadataControl[shared_model.InfoPlist]("/media/mahdi/Cloud/Happle/com.helium.settings/Info.json")
	a, err := infoPlist.Read(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CFBundleDevelopmentRegion: ", a.CFBundleDevelopmentRegion)
}

func testNetwork() {

	userControl := network.NewNetworkManager[[]account.User]("http://localhost:8080/api/v1/user/")

	// Make request (nil body if not needed)
	users, err := userControl.Read("list", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Use the data
	for _, user := range *users {
		fmt.Printf("%d: %s (%s %s)\n",
			user.ID,
			user.Username,
			user.FirstName,
			user.LastName)
	}
}
