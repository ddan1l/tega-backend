package ctx

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type UserContext struct {
	ID       int    `json:"id" example:"1"`
	FullName string `json:"fullName" example:"John"`
	Email    string `json:"email" example:"john@john.com"`
}

func GetUserFromContext(c *gin.Context) (*UserContext, error) {
	user, exists := c.Get("User")

	if !exists {
		return nil, errors.New("user not found in context")
	}

	userObj, ok := user.(UserContext)

	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return &userObj, nil
}
