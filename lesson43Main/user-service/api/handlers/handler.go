package handlers

import (
	_ "github.com/gin-gonic/gin"
	"user_service/strorage/postgres"
)

type HTTPHandler struct {
	UserRepo *postgres.UserRepository
}

func NewHTTPHandler(u *postgres.UserRepository) *HTTPHandler {
	return &HTTPHandler{UserRepo: u}
}
