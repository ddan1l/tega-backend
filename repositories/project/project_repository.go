package project_repository

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
)

type ProjectRepository interface {
	FindByUserId(in *user_dto.FindByIdDto) (*[]models.Project, error)
}
