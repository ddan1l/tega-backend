package auth_usecase

import auth_dto "github.com/ddan1l/tega-api/dto/auth"

type AuthUsecase interface {
	RegisterUser(in *auth_dto.RegisterUserDto) error
}
