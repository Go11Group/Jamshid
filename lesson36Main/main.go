package main

import (
	"github.com/gin-gonic/gin"
	"my_project/controller"
)

func main() {
	r := gin.Default()
	router := controller.RouterController(r)
	err := router.Run(":8087")
	if err != nil {
		panic(err)
	}

}
