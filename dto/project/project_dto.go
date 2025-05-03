package project_dto

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

	FindBySlugDto struct {
		Slug string
	}

	FindBySlugAndUserIdDto struct {
		Slug   string
		UserID int
	}

	ProjectDto struct {
		ID          int    `json:"id" example:"1"`
		Name        string `json:"name" example:"test"`
		Slug        string `json:"slug" example:"test"`
		Description string `json:"description" example:"test description"`
	}

	ProjectsDto struct {
		Projects []ProjectDto
	}
)
