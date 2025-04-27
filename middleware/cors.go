package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// allowedOrigins := []string{
		// 	"https://example.com",
		// 	"http://localhost:3000",
		// }

		// origin := c.Request.Header.Get("Origin")
		// for _, o := range allowedOrigins {
		// 	if origin == o {
		// 		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }
		// log.Println(origin)

		origin := c.Request.Header.Get("Origin")

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
