package project_handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	ProjectsPolicies(c *gin.Context)
	ProjectUser(c *gin.Context)
	ProjectUsers(c *gin.Context)
}
