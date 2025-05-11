package user_handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	User(c *gin.Context)
	UserApp(c *gin.Context)
}
