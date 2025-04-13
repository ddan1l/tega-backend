package factory

import auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"

type UseCaseFactory interface {
	CreateAuthUseCase() auth_usecase.AuthUsecase
}

func (f *DefaultFactory) CreateAuthUseCase() auth_usecase.AuthUsecase {
	return auth_usecase.NewAuthUsecaseImpl(
		f.CreateUserRepository(),
		f.CreateTokenRepository(),
	)
}
