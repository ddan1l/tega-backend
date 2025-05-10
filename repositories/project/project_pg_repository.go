package project_repository

import (
	"errors"

	"github.com/ddan1l/tega-backend/database"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"github.com/ddan1l/tega-backend/models"
	"gorm.io/gorm"
)

type projectPgRepository struct {
	db database.Database
}

func NewProjectPgRepository(db database.Database) ProjectRepository {
	return &projectPgRepository{db: db}
}

func (r *projectPgRepository) WithTx(db database.Database) ProjectRepository {
	return &projectPgRepository{db: db}
}

func (r *projectPgRepository) FindProjectUser(in *project_dto.FindBySlugAndUserIdDto) (*models.ProjectUser, error) {
	var projectUsers []models.ProjectUser

	result := r.db.GetDb().Preload("Project").Preload("Role").Where("user_id = ?", in.UserID).Find(&projectUsers)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, pu := range projectUsers {
		if pu.Project.Slug == in.Slug {
			return &pu, nil
		}
	}

	return nil, errors.New("not fouund")
}

func (r *projectPgRepository) FindProjectsByUserId(in *project_dto.FindByUserIdDto) (*[]models.Project, error) {
	var projectUsers []models.ProjectUser

	result := r.db.GetDb().Preload("Project").Preload("Role").Where("user_id = ?", in.UserID).Find(&projectUsers)

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
