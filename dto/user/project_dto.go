package user_dto

import "github.com/ddan1l/tega-backend/ctx"

type (
	ProjectUserDto struct {
		UserID    int
		RoleID    int
		ProjectID int
	}

	CreateProjectDto struct {
		Project *ProjectDto
		User    *ctx.UserContext
	}

	ProjectDto struct {
		ID          int
		Name        string
		Slug        string
		Description string
	}

	ProjectsDto struct {
		Projects []ProjectDto
	}
)
