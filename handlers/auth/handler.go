package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h *authHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}
