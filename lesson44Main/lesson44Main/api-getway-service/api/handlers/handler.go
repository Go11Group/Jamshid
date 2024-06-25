package handlers

import (
	service "api_get_way/client"
	_ "github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Client *service.UserClient
}

func NewHTTPHandler(u *service.UserClient) *HTTPHandler {
	return &HTTPHandler{Client: u}
}
