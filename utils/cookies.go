package utils

import "github.com/gin-gonic/gin"

func SafeGetCookie(c *gin.Context, name string) string {
	for _, cookie := range c.Request.Cookies() {
		if cookie.Name == name {
			return cookie.Value
		}
	}

	return ""
}
