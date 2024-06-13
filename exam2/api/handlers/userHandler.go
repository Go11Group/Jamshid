package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/strorage/postgres"
	"net/http"
	"time"
)

type HandlerUserConnection struct {
	userHandler *postgres.UserRepository
}

func NewConnectionWithUser(repo *postgres.UserRepository) *HandlerUserConnection {
	return &HandlerUserConnection{userHandler: repo}
}

func (handle *HandlerUserConnection) CreateUserHandler(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.userHandler.CreateUser(user)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}
func (handle *HandlerUserConnection) UpdateUserHandler(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.userHandler.UpdateUser(gn.Param("id"), user)
	if err != nil {
		fmt.Println("+++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}

func (handle *HandlerUserConnection) DeleteUserHandler(gn *gin.Context) {
	err := handle.userHandler.DeleteUser(gn.Param("id"))
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}
func (handle *HandlerUserConnection) GetAllUserHandler(gn *gin.Context) {
	users, err := handle.userHandler.GetAllUsers()
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(users); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":    "Success",
				"status":     http.StatusOK,
				"time":       time.Now(),
				"id":         users[i].Id,
				"name":       users[i].Name,
				"email":      users[i].Email,
				"birthday":   users[i].Birthday,
				"password":   users[i].Password,
				"created_at": users[i].CreatedAt,
				"updated_at": users[i].UpdatedAt,
				"deleted_at": users[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerUserConnection) GetUserFilterHandler(gn *gin.Context) {
	userFilter := model.UserFilter{}
	userFilter.Id = gn.Query("id")
	userFilter.Name = gn.Query("name")
	userFilter.Email = gn.Query("description")
	userFilter.Birthday = gn.Query("birthday")
	userFilter.Password = gn.Query("password")
	users, err := handle.userHandler.GetAllFilter(userFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(users); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":    "Success",
				"status":     http.StatusOK,
				"time":       time.Now(),
				"id":         users[i].Id,
				"name":       users[i].Name,
				"email":      users[i].Email,
				"birthday":   users[i].Birthday,
				"password":   users[i].Password,
				"created_at": users[i].CreatedAt,
				"updated_at": users[i].UpdatedAt,
				"deleted_at": users[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerUserConnection) GetGetCourseByUserIdHandler(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	users, err := handle.userHandler.GetCourseByUser_id(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(users); i++ {
			gn.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"id":      users[i].Id,
				"name":    users[i].Name,
				"email":   users[i].Email,
				"Course": gin.H{
					"id":          users[i].Courses.Id,
					"title":       users[i].Courses.Title,
					"description": users[i].Courses.Description,
					"created_at":  users[i].Courses.CreatedAt,
					"updated_at":  users[i].Courses.UpdatedAt,
					"deleted_at":  users[i].Courses.DeletedAt,
				},
				"birthday":   users[i].Birthday,
				"password":   users[i].Password,
				"created_at": users[i].CreatedAt,
				"updated_at": users[i].UpdatedAt,
				"deleted_at": users[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerUserConnection) GetAllUserByEmailOrNameHandler(gn *gin.Context) {
	users, err := handle.userHandler.GetUserByEmailOrName(gn.Param("name"), gn.Param("email"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(users); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":    "Success",
				"status":     http.StatusOK,
				"time":       time.Now(),
				"id":         users[i].Id,
				"name":       users[i].Name,
				"email":      users[i].Email,
				"birthday":   users[i].Birthday,
				"password":   users[i].Password,
				"created_at": users[i].CreatedAt,
				"updated_at": users[i].UpdatedAt,
				"deleted_at": users[i].DeletedAt,
			})
		}
	}
}
