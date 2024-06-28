package handler

import (
	"context"
	pb "get-way/genproto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBorrow(c *gin.Context) {

	var req pb.CreateBorrowRequest
	if err := c.BindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}

	_, err := h.borrow.CreateBorrow(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	Created(c, err)
}

func (h *Handler) UpdateBorrow(c *gin.Context) {
	var req pb.UpdatedBorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}
	_, err := h.borrow.UpdateBorrow(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	OK(c, err)
}

func (h *Handler) DeleteBorrow(c *gin.Context) {
	var req pb.ByIdRequest
	id := c.Param("id")
	req.Id = id
	_, err := h.borrow.DeleteBorrow(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
	}
	OK(c, err)
}

func (h *Handler) GetBorrow(c *gin.Context) {
	var req pb.BorrowFilterRequest
	req.Userid = c.Query("title")
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
	res, err := h.borrow.GetBorrow(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
