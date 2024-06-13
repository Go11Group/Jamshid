package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

func (handle *Handler) CreateUser(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.UserRepo.CreateUser(user)
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
func (handle *Handler) UpdateUser(gn *gin.Context) {
	user := model.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.UserRepo.UpdateUser(gn.Param("id"), user)
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

func (handle *Handler) DeleteUser(gn *gin.Context) {
	err := handle.UserRepo.DeleteUser(gn.Param("id"))
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

func (handle *Handler) GetUser(gn *gin.Context) {
	userFilter := model.UserFilter{}
	userFilter.Name = gn.Query("name")
	userFilter.Email = gn.Query("description")
	userFilter.Birthday = gn.Query("birthday")
	userFilter.Password = gn.Query("password")
	limit, err := strconv.Atoi(gn.Query("limit"))
	offset, err := strconv.Atoi(gn.Query("offset"))
	userFilter.Limit = limit
	userFilter.Offset = offset
	users, err := handle.UserRepo.GetUser(userFilter)
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

func (handle *Handler) GetCourseByUserId(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	userId, courses, err := handle.UserRepo.GetEnrollmentByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(courses); i++ {
			gn.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"id":      userId,
				"Course": gin.H{
					"id":          courses[i].Id,
					"title":       courses[i].Title,
					"description": courses[i].Description,
				},
			})
		}
	}
}

func (handle *Handler) GetUserByEmailOrName(gn *gin.Context) {
	users, err := handle.UserRepo.GetUserByEmailOrName(gn.Param("name"), gn.Param("email"))
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

func (handle *Handler) GetUserById(gn *gin.Context) {
	user, err := handle.UserRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":    "Success",
			"status":     http.StatusOK,
			"time":       time.Now(),
			"id":         user.Id,
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
		})
	}

}
