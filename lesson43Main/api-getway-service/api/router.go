package api

import (
	"bill_service/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *http.Server {
	router := gin.Default()

	h := handler.NewHandler()

	cr := router.Group("api/card")
	cr.POST("/create", h.BillClient)
	cr.PUT("/update/:id", h.BillClient)
	cr.DELETE("/delete/:id", h.BillClient)
	cr.GET("/get", h.BillClient)
	cr.GET("/id/:id", h.BillClient)
	cr.GET("/card_station/:id", h.BillClient)

	sr := router.Group("api/station")
	sr.POST("/create", h.BillClient)
	sr.PUT("/update/:id", h.BillClient)
	sr.DELETE("/delete/:id", h.BillClient)
	sr.GET("/get", h.BillClient)
	sr.GET("/id/:id", h.BillClient)
	sr.GET("/station_terminal/:id", h.BillClient)

	ter := router.Group("api/terminal")
	ter.POST("/create", h.BillClient)
	ter.PUT("/update/:id", h.BillClient)
	ter.DELETE("/delete/:id", h.BillClient)
	ter.GET("/get", h.BillClient)
	ter.GET("/id/:id", h.BillClient)

	trr := router.Group("api/transaction")
	trr.POST("/create", h.BillClient)
	trr.PUT("/update/:id", h.BillClient)
	trr.DELETE("/delete/:id", h.BillClient)
	trr.GET("/get", h.BillClient)
	trr.GET("/id/:id", h.BillClient)

	us := router.Group("api/user")
	us.POST("/create", h.UserClient)
	us.PUT("/update/:id", h.UserClient)
	us.DELETE("/delete/:id", h.UserClient)
	us.GET("/get", h.UserClient)
	us.GET("/id/:id", h.UserClient)
	us.GET("/user_card_amount/:id", h.UserClient)
	us.GET("/user_card/:id", h.UserClient)
	us.GET("/user_card_credit/:id", h.UserClient)
	us.GET("/user_card_deposit/:id", h.UserClient)

	return &http.Server{Handler: router, Addr: ":8080"}
}
