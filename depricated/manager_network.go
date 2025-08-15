package depricated

import (
	"github.com/mahdi-cpp/api-go-pkg/network"
	"github.com/mahdi-cpp/go-account-service/account"
	"log"
)

type NetworkManager struct {
	networkUser     *network.Control[account.User]
	networkUserList *network.Control[[]account.User]
}

type requestBody struct {
	UserID string `json:"userID"`
}

func NewNetworkAccountManager() *NetworkManager {
	manager := &NetworkManager{
		networkUser:     network.NewNetworkManager[account.User]("http://localhost:8080/api/v1/user/get_user"),
		networkUserList: network.NewNetworkManager[[]account.User]("http://localhost:8080/api/v1/user/list"),
	}

	return manager
}

func (m *NetworkManager) GetUser(id string) (*account.User, error) {

	user, err := m.networkUser.Read("", requestBody{UserID: id})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return user, nil
}

func (m *NetworkManager) GetAll() (*[]account.User, error) {

	users, err := m.networkUserList.Read("", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return users, nil
}

func (m *NetworkManager) GetByFilterOptions(userIDs []string) (*[]account.User, error) {

	users, err := m.networkUserList.Read("", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}
	return users, nil
}
