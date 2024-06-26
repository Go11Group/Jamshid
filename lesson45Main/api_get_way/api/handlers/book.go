package handler

import (
	"context"
	pb "get-way/genproto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBook(c *gin.Context) {

	var req pb.CreateBookRequest
	if err := c.BindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}

	_, err := h.book.CreateBook(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	Created(c, err)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var req pb.UpdatedBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}
	_, err := h.book.UpdateBook(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	OK(c, err)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	var req pb.ByIdRequest
	id := c.Param("id")
	req.Id = id
	_, err := h.book.DeleteBook(context.Background(), &req)
	if err != nil {
		BadRequest(c, err)
		return
	}
	OK(c, err)
}

func (h *Handler) GetBook(c *gin.Context) {
	var req pb.BookFilterRequest
	req.Title = c.Query("title")
	req.Author = c.Query("author")

	published, err := strconv.Atoi(c.Query("published"))
	if err != nil {
		published = 0
	}
	req.Published = int32(published)
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 0
	}
	req.LimitOffset.Limit = int32(limit)
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		limit = 0
	}
	req.LimitOffset.Offset = int32(offset)
	res, err := h.book.GetBook(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
