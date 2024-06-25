package handlers

import (
	pb "api_get_way/proto"
	//"bill_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *HTTPHandler) CreateUser(gn *gin.Context) {
	user := pb.UserCreateRequest{}
	err := gn.BindJSON(&user)
	if err != nil {
		fmt.Println("+++++++++", err)
		BadRequest(gn, err)
		return
	}

	_, err = h.Client.CreateUser(&user)
	if err != nil {
		InternalServerError(gn, err)

		return
	}

	Created(gn, err)
}
func (h *HTTPHandler) UpdateUser(gn *gin.Context) {
	user := pb.UserUpdatedRequest{}
	err := gn.BindJSON(&user)
	if err != nil {
		BadRequest(gn, err)
		return

	}

	_, err = h.Client.UpdateUser(&user)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}

func (h *HTTPHandler) DeleteUser(gn *gin.Context) {
	id := pb.GetByIdRequest{
		Id: gn.Param("id"),
	}
	_, err := h.Client.DeleteUser(id)
	if err != nil {
		fmt.Println("_____", err)
		InternalServerError(gn, err)
		return

	}
	OK(gn, err)

}

func (h *HTTPHandler) GetUser(gn *gin.Context) {
	userFilter := pb.UserFilterRequest{}
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
	userFilter.Age = int32(age)
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
	}
	userFilter.Limit = int32(limit)
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)

	}
	userFilter.Offset = int32(offset)
	users, err := h.Client.GetUser(&userFilter)
	if err != nil {
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"users": users,
	})
}

//func (h *HTTPHandler) GetUserById(gn *gin.Context) {
//
//	if err != nil {
//		fmt.Println("++++++", err)
//		InternalServerError(gn, err)
//
//	}
//	gn.JSON(200, gin.H{
//		"users": users,
//	})
//}
//func (h *HTTPHandler) GetAmountUser(gn *gin.Context) {
//	id:=pb.GetByIdRequest{
//		Id: gn.Param("id"),
//	}
//	h.Client.
//	usersAmount, err := h.Client.
//	if err != nil {
//		fmt.Println("++++++", err)
//		InternalServerError(gn, err)
//
//	}
//	gn.JSON(200, gin.H{
//		"usersAmount": usersAmount,
//	})
//}

func (h *HTTPHandler) GetUserCard(gn *gin.Context) {
	id := pb.GetByIdRequest{Id: gn.Param("id")}
	userCards, err := h.Client.GetUserCard(&id)
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCards": userCards,
	})
}
func (h *HTTPHandler) GetUserCardDeposit(gn *gin.Context) {
	id := pb.GetByIdRequest{Id: gn.Param("id")}
	userCardDeposits, err := h.Client.GetUserDepositCard(&id)
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCardDeposits": userCardDeposits,
	})
}
func (h *HTTPHandler) GetUserCardCredit(gn *gin.Context) {
	id := pb.GetByIdRequest{Id: gn.Param("id")}
	userCardCredits, err := h.Client.GetUserCreditCard(&id)
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"userCardCredits": userCardCredits,
	})
}
