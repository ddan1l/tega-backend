package auth_middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Middleware() gin.HandlerFunc
}
