package api

import (
	"bill_service/api/handler"
	"bill_service/storage/postgres"
	"github.com/gin-gonic/gin"
)

func RouterApi(c *postgres.CardRepository, s *postgres.StationRepository, t *postgres.TerminalRepository, tr *postgres.TransactionRepository) *gin.Engine {
	routerHandle := handler.NewHTTPHandler(c, s, t, tr)
	router := gin.Default()
	cr := router.Group("api/card")
	cr.POST("/create", routerHandle.CreateCard)
	cr.PUT("/update/:id", routerHandle.UpdateCard)
	cr.DELETE("/delete/:id", routerHandle.DeleteCard)
	cr.GET("/get", routerHandle.GetCard)
	cr.GET("/id/:id", routerHandle.GetCardById)
	cr.GET("/card_station/:id", routerHandle.GetCardById)

	sr := router.Group("api/station")
	sr.POST("/create", routerHandle.CreateStation)
	sr.PUT("/update/:id", routerHandle.UpdateStation)
	sr.DELETE("/delete/:id", routerHandle.DeleteStation)
	sr.GET("/get", routerHandle.GetStation)
	sr.GET("/id/:id", routerHandle.GetStationById)
	sr.GET("/station_terminal/:id", routerHandle.GetCardWithTransactionStations)

	ter := router.Group("api/terminal")
	ter.POST("/create", routerHandle.CreateTerminal)
	ter.PUT("/update/:id", routerHandle.UpdateTerminal)
	ter.DELETE("/delete/:id", routerHandle.DeleteTerminal)
	ter.GET("/get", routerHandle.GetTerminal)
	ter.GET("/id/:id", routerHandle.GetTerminalById)

	trr := router.Group("api/transaction")
	trr.POST("/create", routerHandle.CreateTransaction)
	trr.PUT("/update/:id", routerHandle.UpdateTransaction)
	trr.DELETE("/delete/:id", routerHandle.DeleteTransaction)
	trr.GET("/get", routerHandle.GetTransaction)
	trr.GET("/id/:id", routerHandle.GetTransactionById)
	return router

}
