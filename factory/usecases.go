package factory

import (
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
	user_usercase "github.com/ddan1l/tega-backend/usecases/user"
)

type UseCaseFactory interface {
	CreateAuthUseCase() auth_usecase.AuthUsecase
	CreateUserUseCase() user_usercase.UserUsecase
}

func (f *DefaultFactory) CreateAuthUseCase() auth_usecase.AuthUsecase {
	return auth_usecase.NewAuthUsecaseImpl(
		f.CreateUserRepository(),
		f.CreateTokenRepository(),
	)
}

func (f *DefaultFactory) CreateUserhUseCase() user_usercase.UserUsecase {
	return user_usercase.NewUserUsecaseImpl(
		f.CreateUserRepository(),
		f.CreateProjectRepository(),
	)
}
