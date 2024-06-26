package handler

import (
	"context"
	pb "get-way/genproto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(c *gin.Context) {

	var req pb.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}

	_, err := h.user.CreateUser(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	Created(c, err)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req pb.UpdatedUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}
	_, err := h.user.UpdateUser(context.Background(), &req)
	if err != nil {
		InternalServerError(c, err)
		return
	}
	OK(c, err)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var req pb.ByIdRequest
	id := c.Param("id")
	req.Id = id
	_, err := h.user.DeleteUser(context.Background(), &req)
	if err != nil {
		BadRequest(c, err)
		return
	}
	OK(c, err)
}

func (h *Handler) GetUser(c *gin.Context) {
	var req pb.UserFilterRequest
	req.Name = c.Query("first_name")
	req.LastName = c.Query("last_name")
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

	res, err := h.user.GetUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
