package api

import (
	handler "api_getway/api/handlers"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func Router(p *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandlerStruct(p)
	user := router.Group("/transport")
	{
		user.POST("/report_traffic_jam/bus_number", h.ReportTrafficJam)
		user.PUT("/trac_location/bus_number", h.TrackBusLocation)
		user.DELETE("/bus_schedule", h.GetBusSchedule)
	}

	book := router.Group("/weather")
	{
		book.POST("/report_weather", h.ReportWeatherCondition)
		book.PUT("/weather_forecast", h.GetWeatherForecast)
		book.DELETE("/current_weather", h.GetCurrentWeather)
	}

	return router

}
