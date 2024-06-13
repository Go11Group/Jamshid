package main

import (
	"github.com/gin-gonic/gin"
	"my_project/controller"
)

func main() {
	r := gin.Default()
	router := controller.Router(r)
	err := router.Run(":8090")
	if err != nil {
		return
	}
}
