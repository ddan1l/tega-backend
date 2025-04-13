package auth_usecase

import (
	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	errs "github.com/ddan1l/tega-backend/errors"
)

type AuthUsecase interface {
	RegisterUser(in *auth_dto.RegisterUserDto) (*auth_dto.TokensPairDto, *errs.AppError)
	CheckUserExists(email string) *errs.AppError
}
