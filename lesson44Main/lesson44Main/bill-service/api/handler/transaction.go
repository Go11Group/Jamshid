package handler

import (
	"bill_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *HTTPHandler) CreateTransaction(gn *gin.Context) {

	transaction := models.Transaction{}
	err := gn.BindJSON(&transaction)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	if transaction.TransactionType == "deposit" {
		transaction.Amount = (-1) * transaction.Amount
	}

	tran, err := h.TransactionRepo.CreateTransaction(&transaction)
	if tran == false {
		gn.JSON(200, gin.H{
			"message": "Hisobda yetarli mablag' yoq",
		})
	}

	if err != nil {
		InternalServerError(gn, err)

		return
	}
	Created(gn, err)

}
func (h *HTTPHandler) UpdateTransaction(gn *gin.Context) {

	transaction := models.Transaction{}
	err := gn.BindJSON(&transaction)
	if err != nil {
		BadRequest(gn, err)

		return
	}

	err = h.TransactionRepo.UpdateTransaction(gn.Param("id"), &transaction)
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	OK(gn, err)
}
func (h *HTTPHandler) DeleteTransaction(gn *gin.Context) {
	err := h.TransactionRepo.DeletedTransaction(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	OK(gn, err)

}

func (h *HTTPHandler) GetTransaction(gn *gin.Context) {

	transactionFilter := models.Filter{}

	transactionFilter.UserId = gn.Query("user_id")
	transactionFilter.TerminalId = gn.Query("terminal_id")
	transactionFilter.TransactionType = gn.Query("transaction_type")
	amount1 := gn.Query("amount")
	if len(amount1) == 0 {
		amount1 = "0"
	}
	amount, err := strconv.Atoi(amount1)
	if err != nil {
		BadRequest(gn, err)

	}
	transactionFilter.Amount = amount
	fmt.Println("-------", transactionFilter.Amount)
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)

	if err != nil {

		BadRequest(gn, err)

		return
	}
	transactionFilter.Limit = limit
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {

		BadRequest(gn, err)

		return
	}
	transactionFilter.Offset = offset
	transactions, err := h.TransactionRepo.GetTransaction(transactionFilter)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"transactions": transactions,
	})
}

func (h *HTTPHandler) GetTransactionById(gn *gin.Context) {

	transaction, err := h.TransactionRepo.GetTransactionById(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	gn.JSON(200, gin.H{
		"transaction": transaction,
	})
}
