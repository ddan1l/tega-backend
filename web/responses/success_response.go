package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
