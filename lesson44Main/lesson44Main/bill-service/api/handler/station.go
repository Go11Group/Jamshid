package handler

import (
	"bill_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
)

func (h *HTTPHandler) CreateStation(gn *gin.Context) {

	station := models.Station{}
	err := gn.BindJSON(&station)
	if err != nil {
		BadRequest(gn, err)

		return
	}
	err = h.StationRepo.CreateStation(&station)
	if err != nil {
		slog.Info("message", err.Error())
		fmt.Println("__________+++++++++++++++++++", err)
		slog.Error("message", err.Error())
		InternalServerError(gn, err)
		return
	}

	Created(gn, err)
}
func (h *HTTPHandler) UpdateStation(gn *gin.Context) {
	station := models.Station{}
	err := gn.BindJSON(&station)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	err = h.StationRepo.UpdateStation(gn.Param("id"), &station)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}
func (h *HTTPHandler) DeleteStation(gn *gin.Context) {
	err := h.StationRepo.DeletedStation(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)
}

func (h *HTTPHandler) GetStation(gn *gin.Context) {
	stationFilter := models.Filter{}

	stationFilter.Name = gn.Query("name")
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	stationFilter.Limit = limit
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	stationFilter.Offset = offset
	stations, err := h.StationRepo.GetStation(stationFilter)
	if err != nil {
		fmt.Println("errror++++++++", err)
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"stations": stations,
	})
}

func (h *HTTPHandler) GetStationById(gn *gin.Context) {
	station, err := h.StationRepo.GetStationById(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"station": station,
	})
}

func (h *HTTPHandler) StationTerminal(gn *gin.Context) {
	stationTerminal, err := h.StationRepo.GetTerminalsByStationId(gn.Param("id"))
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, gin.H{
		"station": stationTerminal,
	})
}
