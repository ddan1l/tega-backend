package project_handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	ProjectsPolicies(c *gin.Context)
}
