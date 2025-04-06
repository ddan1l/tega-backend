package user_repository

import user_dto "github.com/ddan1l/tega-api/dto/user"

type UserRepository interface {
	Create(in *user_dto.CreateUserDto) error
	FindById(in *user_dto.FindByIdDto) error
}
