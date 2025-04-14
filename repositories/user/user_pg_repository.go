package user_repository

import (
	"errors"

	"github.com/ddan1l/tega-backend/database"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
	"gorm.io/gorm"
)

type userPgRepository struct {
	db database.Database
}

func NewUserPgRepository(db database.Database) UserRepository {
	return &userPgRepository{db: db}
}

func (r *userPgRepository) Create(in *user_dto.CreateUserDto) (*models.User, error) {
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

func (r *userPgRepository) FindById(in *user_dto.FindByIdDto) (*models.User, error) {
	var user models.User

	result := r.db.GetDb().Where(models.User{
		ID: in.ID,
	}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userPgRepository) FindByEmail(in *user_dto.FindByEmailDto) (*models.User, error) {
	var user models.User

	result := r.db.GetDb().Where(models.User{
		Email: in.Email,
	}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &user, nil
}
