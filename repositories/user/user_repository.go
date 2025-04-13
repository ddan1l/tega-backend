package user_repository

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
)

type UserRepository interface {
	Create(in *user_dto.CreateUserDto) (*models.User, error)
	FindById(in *user_dto.FindByIdDto) (*models.User, error)
	FindByEmail(in *user_dto.FindByEmailDto) (*models.User, error)
}
