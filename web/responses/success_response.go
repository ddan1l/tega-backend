package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
}

type SuccessWithDataResponse struct {
	Success bool `json:"success" example:"true"`
	Data    any  `json:"data"`
}

func Succes(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{Success: true})
}

func SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessWithDataResponse{
		Success: true,
		Data:    data,
	})
}
