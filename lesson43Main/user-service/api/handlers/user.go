package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"user_service/models"
)

func (h *HTTPHandler) CreateUser(gn *gin.Context) {
	user := models.User{}
	err := gn.BindJSON(&user)
	if err != nil {
		fmt.Println("+++++++++", err)
		BadRequest(gn, err)
		return
	}

	err = h.UserRepo.CreateUser(&user)
	if err != nil {
		InternalServerError(gn, err)

		return
	}

	Created(gn, err)
}
func (h *HTTPHandler) UpdateUser(gn *gin.Context) {
	user := models.User{}
	err := gn.BindJSON(&user)
	if err != nil {

		BadRequest(gn, err)
		return

	}

	err = h.UserRepo.UpdateUser(gn.Param("id"), &user)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}

func (h *HTTPHandler) DeleteUser(gn *gin.Context) {
	err := h.UserRepo.DeleteUser(gn.Param("id"))
	if err != nil {
		fmt.Println("_____", err)
		InternalServerError(gn, err)
		return

	}
	OK(gn, err)

}

func (h *HTTPHandler) GetUser(gn *gin.Context) {
	userFilter := models.Filter{}
	userFilter.Name = gn.Query("name")
	userFilter.Phone = gn.Query("phone")
	age1 := gn.Query("age")
	if len(age1) == 0 {
		age1 = "0"
	}
	age, err := strconv.Atoi(age1)
	if err != nil {
		BadRequest(gn, err)

	}
	userFilter.Age = age
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
	}
	userFilter.Limit = limit
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)

	}
	userFilter.Offset = offset
	users, err := h.UserRepo.GetUser(userFilter)
	if err != nil {
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"users": users,
	})
}

func (h *HTTPHandler) GetUserById(gn *gin.Context) {
	users, err := h.UserRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"users": users,
	})
}
func (h *HTTPHandler) GetAmountUser(gn *gin.Context) {
	usersAmount, err := h.UserRepo.GetUserCardAmount(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"usersAmount": usersAmount,
	})
}

func (h *HTTPHandler) GetUserCard(gn *gin.Context) {
	userCards, err := h.UserRepo.GetUserCard(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCards": userCards,
	})
}
func (h *HTTPHandler) GetUserCardDeposit(gn *gin.Context) {
	userCardDeposits, err := h.UserRepo.GetUserDepositCard(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCardDeposits": userCardDeposits,
	})
}
func (h *HTTPHandler) GetUserCardCredit(gn *gin.Context) {
	userCardCredits, err := h.UserRepo.GetUserCreditCard(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCardCredits": userCardCredits,
	})
}
