package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
}

type SuccessWithDataResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}

func Succes(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{Success: true})
}

func SuccessWithData[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, SuccessWithDataResponse[T]{
		Success: true,
		Data:    data,
	})
}
