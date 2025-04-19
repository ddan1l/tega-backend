package res

import (
	errs "github.com/ddan1l/tega-backend/errors"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool           `json:"success" example:"false"`
	Error   *errs.AppError `json:"error"`
}

func Error(c *gin.Context, err *errs.AppError) {
	c.JSON(err.Status, ErrorResponse{Success: false, Error: err})
}
