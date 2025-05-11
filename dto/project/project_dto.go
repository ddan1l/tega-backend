package project_dto

import user_dto "github.com/ddan1l/tega-backend/dto/user"

type (
	ProjectUserDto struct {
		ID        int               `json:"id" example:"1"`
		UserID    int               `json:"user_id" example:"1"`
		RoleID    int               `json:"role_id" example:"1"`
		ProjectID int               `json:"project_id" example:"1"`
		Project   *ProjectDto       `json:"project" example:"1"`
		User      *user_dto.UserDto `json:"user" example:"1"`
	}

	CreateProjectDto struct {
		Project *ProjectDto
		UserID  int
	}

	FindBySlugDto struct {
		Slug string
	}

	FindByIdDto struct {
		ID int
	}

	FindBySlugAndUserIdDto struct {
		Slug   string
		UserID int
	}

	FindByUserIdDto struct {
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
