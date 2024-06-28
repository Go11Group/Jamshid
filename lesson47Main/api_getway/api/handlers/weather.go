package handler

import (
	pb "api_getway/genproto"
	"context"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(c *gin.Context) {

	var req pb.CurrentWheatherRequest
	place := c.Query("place")
	req.City = place

	_, err := h.weather.GetCurrentWeather(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	Created(c, err)
}

func (h *Handler) GetWeatherForecast(c *gin.Context) {
	var req pb.ForecastWheatherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}
	resp, err := h.weather.GetWeatherForecast(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	c.JSON(200, resp)
}

func (h *Handler) ReportWeatherCondition(c *gin.Context) {
	var req pb.ReportWheatherRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		BadRequest(c, err)
	}
	resp, err := h.weather.ReportWeatherCondition(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
	}
	c.JSON(200, resp)
}
