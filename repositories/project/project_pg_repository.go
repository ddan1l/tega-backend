package project_repository

import (
	"github.com/ddan1l/tega-backend/database"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
)

type projectPgRepository struct {
	db database.Database
}

func NewProjectPgRepository(db database.Database) ProjectRepository {
	return &projectPgRepository{db: db}
}

func (r *projectPgRepository) Create(in *user_dto.CreateUserDto) (*models.User, error) {
	user := &models.User{
		FullName:     in.FullName,
		Email:        in.Email,
		PasswordHash: in.PasswordHash,
	}

	result := r.db.GetDb().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *projectPgRepository) FindByUserId(in *user_dto.FindByIdDto) (*[]models.Project, error) {
	var projectUsers []models.ProjectUsers

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
