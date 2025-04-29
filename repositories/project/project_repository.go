package project_repository

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
)

type ProjectRepository interface {
	FindProjectsByUserId(in *user_dto.FindByIdDto) (*[]models.Project, error)
	CreateProject(in *user_dto.ProjectDto) (*models.Project, error)
	CreateProjectUser(in *user_dto.ProjectUserDto) (*models.ProjectUser, error)
}
