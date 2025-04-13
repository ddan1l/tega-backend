package auth_handler

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register(c *gin.Context)
	// Register(c any) error
	// Logout(c any) error
}
