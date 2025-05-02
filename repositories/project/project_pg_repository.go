package project_repository

import (
	"errors"

	"github.com/ddan1l/tega-backend/database"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
	"gorm.io/gorm"
)

type projectPgRepository struct {
	db database.Database
}

func NewProjectPgRepository(db database.Database) ProjectRepository {
	return &projectPgRepository{db: db}
}

func (r *projectPgRepository) FindProjectsByUserId(in *user_dto.FindByIdDto) (*[]models.Project, error) {
	var projectUsers []models.ProjectUser

	result := r.db.GetDb().Preload("Project").Where("user_id = ?", in.ID).Find(&projectUsers)

	if result.Error != nil {
		return nil, result.Error
	}

	projects := make([]models.Project, 0, len(projectUsers))

	for _, pu := range projectUsers {
		projects = append(projects, pu.Project)
	}

	return &projects, nil
}

func (r *projectPgRepository) FindProjectsBySlug(in *project_dto.FindBySlugDto) (*models.Project, error) {
	var project models.Project

	result := r.db.GetDb().Where(models.Project{
		Slug: in.Slug,
	}).First(&project)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &project, nil
}

func (r *projectPgRepository) CreateProject(in *project_dto.ProjectDto) (*models.Project, error) {
	project := &models.Project{
		Name:        in.Name,
		Slug:        in.Slug,
		Description: in.Description,
	}

	result := r.db.GetDb().Create(&project)

	if result.Error != nil {
		return nil, result.Error
	}

	return project, nil
}

func (r *projectPgRepository) CreateProjectUser(in *project_dto.ProjectUserDto) (*models.ProjectUser, error) {
	projectUser := &models.ProjectUser{
		RoleID:    in.RoleID,
		UserID:    in.UserID,
		ProjectID: in.ProjectID,
	}

	result := r.db.GetDb().Create(&projectUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return projectUser, nil
}
