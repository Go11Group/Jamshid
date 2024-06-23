package api

import (
	"github.com/gin-gonic/gin"
	"user_service/api/handlers"
	"user_service/strorage/postgres"
)

func RouterApi(ur *postgres.UserRepository) *gin.Engine {
	router := gin.Default()
	routerHandle := handlers.NewHTTPHandler(ur)
	us := router.Group("api/user")
	us.POST("/create", routerHandle.CreateUser)
	us.PUT("/update/:id", routerHandle.UpdateUser)
	us.DELETE("/delete/:id", routerHandle.DeleteUser)
	us.GET("/get", routerHandle.GetUser)
	us.GET("/id/:id", routerHandle.GetUserById)
	us.GET("/user_card_amount/:id", routerHandle.GetAmountUser)
	us.GET("/user_card/:id", routerHandle.GetUserCard)
	us.GET("/user_card_credit/:id", routerHandle.GetUserCardCredit)
	us.GET("/user_card_deposit/:id", routerHandle.GetUserCardDeposit)
	return router
}
