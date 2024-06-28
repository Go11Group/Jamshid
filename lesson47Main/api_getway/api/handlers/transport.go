package handler

import (
	pb "api_getway/genproto"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetBusSchedule(c *gin.Context) {
	var req pb.BusScheduleRequest
	number, err := strconv.Atoi(c.Param("bus_number"))
	if err != nil {
		BadRequest(c, err)
	}
	req.BusNumber = int32(number)

	res, err := h.transport.GetBusSchedule(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(200, res)
}

func (h *Handler) TrackBusLocation(c *gin.Context) {
	var req pb.BusLocationRequest
	number, err := strconv.Atoi(c.Param("bus_number"))
	if err != nil {
		BadRequest(c, err)
	}
	req.BusNumber = int32(number)
	resp, err := h.transport.TrackBusLocation(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	c.JSON(200, resp)
}

func (h *Handler) ReportTrafficJam(c *gin.Context) {
	var req *pb.TrafficJamRequest
	report := c.Query("report")
	req.Report = report
	_, err := h.transport.ReportTrafficJam(context.Background(), &req)
	if err != nil {
		BadRequest(c, err)
		return
	}
	OK(c, err)
}
