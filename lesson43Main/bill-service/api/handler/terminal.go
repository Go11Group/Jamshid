package handler

import (
	"bill_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *HTTPHandler) CreateTerminal(gn *gin.Context) {

	terminal := models.Terminal{}
	err := gn.BindJSON(&terminal)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	err = h.TerminalRepo.CreateTerminal(&terminal)
	if err != nil {
		fmt.Println("_____", err)
		InternalServerError(gn, err)
		return
	}
	Created(gn, err)
}
func (h *HTTPHandler) UpdateTerminal(gn *gin.Context) {

	terminal := models.Terminal{}
	err := gn.BindJSON(&terminal)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	err = h.TerminalRepo.UpdateTerminal(gn.Param("id"), &terminal)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}
func (h *HTTPHandler) DeleteTerminal(gn *gin.Context) {
	err := h.TerminalRepo.DeletedTerminal(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}

func (h *HTTPHandler) GetTerminal(gn *gin.Context) {
	terminalFilter := models.Filter{}
	terminalFilter.StationId = gn.Query("station_id")
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	terminalFilter.Limit = limit
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	terminalFilter.Offset = offset
	terminals, err := h.TerminalRepo.GetTerminal(terminalFilter)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"terminals": terminals,
	})
}

func (h *HTTPHandler) GetTerminalById(gn *gin.Context) {
	terminal, err := h.TerminalRepo.GetTerminalById(gn.Param("id"))
	if err != nil {
		fmt.Println("++++++++", err)
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"terminal": terminal,
	})
}
