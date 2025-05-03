package ctx

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ProjectContext struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"test"`
	Slug        string `json:"slug" example:"test"`
	Description string `json:"description" example:"test description"`
}

func GetProjectFromContext(c *gin.Context) (*ProjectContext, error) {
	project, exists := c.Get("Project")

	if !exists {
		return nil, errors.New("project not found in context")
	}

	projectObj, ok := project.(ProjectContext)

	if !ok {
		return nil, errors.New("invalid project type in context")
	}

	return &projectObj, nil
}
