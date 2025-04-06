package user_repository

import (
	"fmt"

	"github.com/ddan1l/tega-api/database"
	user_dto "github.com/ddan1l/tega-api/dto/user"
	"github.com/ddan1l/tega-api/models"
)

type userPgRepository struct {
	db database.Database
}

func NewUserPgRepository(db database.Database) UserRepository {
	return &userPgRepository{db: db}
}

func (r *userPgRepository) Create(in *user_dto.CreateUserDto) error {
	data := &models.User{
		FullName:     in.FullName,
		Email:        in.Email,
		PasswordHash: in.PasswordHash,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		fmt.Printf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	fmt.Printf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *userPgRepository) FindById(in *user_dto.FindByIdDto) error {
	result := r.db.GetDb().First(&models.User{}, in.Id)

	if result.Error != nil {
		fmt.Printf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	fmt.Printf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
