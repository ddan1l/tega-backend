package auth_usecase

import (
	"time"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	errs "github.com/ddan1l/tega-backend/errors"
)

const (
	AccessTokenMaxAge  = time.Minute * 15
	RefreshTokenMaxAge = time.Hour * 24 * 14
)

type AuthUsecase interface {
	RegisterUser(in *auth_dto.RegisterUserDto) (*auth_dto.TokensPairDto, *errs.AppError)
	LoginUser(in *auth_dto.LoginUserDto) (*auth_dto.TokensPairDto, *errs.AppError)
	CheckUserExists(email string) *errs.AppError
	Authenticate(in *auth_dto.TokensPairDto) (*auth_dto.AuthenticatedDto, *errs.AppError)
	DeleteToken(t string) *errs.AppError
}
