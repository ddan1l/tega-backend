package res

import (
	errs "github.com/ddan1l/tega-backend/errors"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err *errs.AppError) {
	c.JSON(err.Status, gin.H{"error": &err, "success": false})
}
