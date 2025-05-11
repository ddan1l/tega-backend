package ctx

import (
	"errors"

	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (*user_dto.UserDto, error) {
	user, exists := c.Get("User")

	if !exists {
		return nil, errors.New("user not found in context")
	}

	userObj, ok := user.(user_dto.UserDto)

	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return &userObj, nil
}
