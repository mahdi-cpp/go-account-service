package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	manager *ServiceManager
}

func NewUserHandler(manager *ServiceManager) *UserHandler {
	return &UserHandler{
		manager: manager,
	}
}

type requestBody struct {
	UserID string `json:"userID"`
}

func (handler *UserHandler) Create(c *gin.Context) {

	userID, err := GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(userID)

	var request Update
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//userStorage, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}

	newItem, err := handler.manager.Create(&User{
		Username:    request.Username,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		AvatarURL:   request.AvatarURL,
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//update := asset.AssetUpdate{AssetIds: request.AssetIds, AddAlbums: []int{newItem.ID}}
	//_, err = userStorage.UpdateAsset(update)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//userStorage.UpdateCollections()

	c.JSON(http.StatusCreated, newItem)
}

func (handler *UserHandler) Update(c *gin.Context) {

	//userID, err := utils.GetUserId(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	var itemUpdate Update

	if err := c.ShouldBindJSON(&itemUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update, err := handler.manager.Update(itemUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, update)
}

func (handler *UserHandler) Delete(c *gin.Context) {

	userID, err := GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//var item account.User
	//if err := c.ShouldBindJSON(&item); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}

	err = handler.manager.UserCollection.Delete(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *UserHandler) GetCollectionList(c *gin.Context) {

	item2, err := handler.manager.UserCollection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *UserHandler) GetUserByID(c *gin.Context) {

	userID, err := GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(userID)

	var request requestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := handler.manager.UserCollection.Get(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//result := asset.PHCollectionList[*account.User]{
	//	Collections: make([]*asset.PHCollection[*account.User], len(items)),
	//}
	//
	//for i, item := range items {
	//	assets, _ := handler.manager.AccountManager.GetItemAssets(item.ID)
	//	result.Collections[i] = &asset.PHCollection[*account.User]{
	//		Item:   item,
	//		Assets: assets,
	//	}
	//}

	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) GetUser(c *gin.Context) {

	item, err := handler.manager.UserCollection.Get("0198adfd-c0ca-7151-990f-b50956fc7f27")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (handler *UserHandler) GetList(c *gin.Context) {

	//var with asset.PHFetchOptions
	//if err := c.ShouldBindJSON(&with); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	fmt.Println("Invalid request")
	//	return
	//}

	//userStorage, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}

	//items, err := handler.manager.AccountManager.GetAllSorted(with.SortBy, with.SortOrder)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	items, err := handler.manager.UserCollection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//result := asset.PHCollectionList[*account.User]{
	//	Collections: make([]*asset.PHCollection[*account.User], len(items)),
	//}
	//
	//for i, item := range items {
	//	assets, _ := handler.manager.AccountManager.GetItemAssets(item.ID)
	//	result.Collections[i] = &asset.PHCollection[*account.User]{
	//		Item:   item,
	//		Assets: assets,
	//	}
	//}

	c.JSON(http.StatusOK, items)
}
