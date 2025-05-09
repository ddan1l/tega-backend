package project_repository

import (
	"github.com/ddan1l/tega-backend/database"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
)

type ProjectRepository interface {
	FindProjectsByUserId(in *user_dto.FindByIdDto) (*[]models.Project, error)
	FindProjectsBySlug(in *project_dto.FindBySlugDto) (*models.Project, error)
	CreateProject(in *project_dto.ProjectDto) (*models.Project, error)
	CreateProjectUser(in *project_dto.ProjectUserDto) (*models.ProjectUser, error)
	WithTx(db database.Database) ProjectRepository
}
