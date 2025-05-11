package ctx

import (
	"errors"

	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"github.com/gin-gonic/gin"
)

func GetProjectUserFromContext(c *gin.Context) (*project_dto.ProjectUserDto, error) {
	projectUser, exists := c.Get("ProjectUser")

	if !exists {
		return nil, errors.New("project not found in context")
	}

	projectUserObj, ok := projectUser.(project_dto.ProjectUserDto)

	if !ok {
		return nil, errors.New("invalid project user type in context")
	}

	return &projectUserObj, nil
}
