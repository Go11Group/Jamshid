package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithUserHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewUserRepository(db)

}

type ConnectionWithUser struct {
	userHandler *postgres.UserRepository
}

func NewConnectionWithUser(repo *postgres.UserRepository) *ConnectionWithUser {
	return &ConnectionWithUser{userHandler: repo}
}

func (ush *ConnectionWithUser) CreateUserHandler(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
	err = ush.userHandler.CreateUser(user)
	if err != nil {
		fmt.Println("------------", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "created is success",
		"status":  http.StatusOK,
	})
}
func (ush *ConnectionWithUser) UpdateUserHandler(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}

	err = ush.userHandler.UpdateUser(gn.Param("id"), user)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "updated is success",
		"status":  http.StatusOK,
	})
}
func (ush *ConnectionWithUser) DeletedUserHandler(gn *gin.Context) {
	err := ush.userHandler.DeletedUser(gn.Param("id"))
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "deleted is success",
		"status":  http.StatusOK,
	})
}
func (ush *ConnectionWithUser) GetAllUserHandler(gn *gin.Context) {
	users := []model.User{}
	users, err := ush.userHandler.GetAllUser()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	for i := 0; i < len(users); i++ {
		gn.JSON(http.StatusOK, gin.H{
			"id":           users[i].Id,
			"name":         users[i].Name,
			"email":        users[i].Email,
			"phone_number": users[i].PhoneNumber,
			"birthday":     users[i].Birthday,
			"gender":       users[i].Gender,
			"created_at":   users[i].CreatedAt,
			"updated_at":   users[i].UpdatedAt,
			"deleted_at":   users[i].DeletedAt,
		})
	}

}
