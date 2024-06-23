package handler

import (
	"bill_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
)

func (h *HTTPHandler) CreateCard(gn *gin.Context) {
	card := models.Card{}
	slog.Info("Create Card start")
	err := gn.BindJSON(&card)
	if err != nil {
		slog.Error("Error", err)
		BadRequest(gn, err)
		return
	}
	slog.Info("Server start created card ")
	err = h.CardRepo.CreateCard(&card)
	if err != nil {
		fmt.Println("-----------", err)
		slog.Info("error", err)

		InternalServerError(gn, err)
		return
	}
	slog.Info("Status ", 200)
	Created(gn, err)

}
func (h *HTTPHandler) UpdateCard(gn *gin.Context) {

	card := models.Card{}
	err := gn.BindJSON(&card)
	if err != nil {
		BadRequest(gn, err)

		return
	}

	err = h.CardRepo.UpdateCard(gn.Param("id"), &card)
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	OK(gn, err)

}
func (h *HTTPHandler) DeleteCard(gn *gin.Context) {
	err := h.CardRepo.DeletedCard(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	OK(gn, err)
}

func (h *HTTPHandler) GetCard(gn *gin.Context) {

	cardFilter := models.Filter{}
	cardFilter.Number = gn.Query("number")
	cardFilter.UserId = gn.Query("user_id")
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)

		return
	}
	cardFilter.Limit = limit
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)

		return
	}
	cardFilter.Offset = offset
	cards, err := h.CardRepo.GetCard(cardFilter)
	if err != nil {
		fmt.Println("++++++", err)
		BadRequest(gn, err)

		return
	}
	gn.JSON(200, gin.H{
		"cards": cards,
	})
}

func (h *HTTPHandler) GetCardById(gn *gin.Context) {

	cards, err := h.CardRepo.GetCardById(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	gn.JSON(200, gin.H{
		"cards": cards,
	})
}
func (h *HTTPHandler) GetCardAmount(gn *gin.Context) {

	cards, err := h.CardRepo.GetCardById(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)

		return
	}
	gn.JSON(200, gin.H{
		"cards": cards,
	})
}
func (h *HTTPHandler) GetCardWithTransactionStations(gn *gin.Context) {

	cardStations, err := h.CardRepo.GetStationByCardId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(gn, err)

		return
	}
	gn.JSON(200, gin.H{
		"cardStations": cardStations,
	})
}
