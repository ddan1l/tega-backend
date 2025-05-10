package project_repository

import (
	"github.com/ddan1l/tega-backend/database"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"github.com/ddan1l/tega-backend/models"
)

type ProjectRepository interface {
	FindProjectsByUserId(in *project_dto.FindByUserIdDto) (*[]models.Project, error)
	FindProjectsBySlug(in *project_dto.FindBySlugDto) (*models.Project, error)
	FindProjectUser(in *project_dto.FindBySlugAndUserIdDto) (*models.ProjectUser, error)
	CreateProject(in *project_dto.ProjectDto) (*models.Project, error)
	CreateProjectUser(in *project_dto.ProjectUserDto) (*models.ProjectUser, error)
	WithTx(db database.Database) ProjectRepository
}
