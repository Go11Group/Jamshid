package api

import (
	"get-way/api/handlers"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func Router(p *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandlerStruct(p)
	user := router.Group("/user")
	{
		user.POST("/create", h.CreateUser)
		user.PUT("/update", h.UpdateUser)
		user.DELETE("/delete/:id", h.DeleteUser)
		user.GET("/get", h.GetUser)
	}

	book := router.Group("/book")
	{
		book.POST("/create", h.CreateBook)
		book.PUT("/update", h.UpdateBook)
		book.DELETE("/delete/:id", h.DeleteBook)
		book.GET("/get", h.GetBook)
	}

	borrow := router.Group("/borrow")
	{
		borrow.POST("/create", h.CreateBorrow)
		borrow.PUT("/update", h.UpdateBorrow)
		borrow.DELETE("/delete/:id", h.DeleteBorrow)
		borrow.GET("/getallbyuser", h.GetBorrow)
	}

	return router

}
