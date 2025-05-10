package factory

import (
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
	project_usecase "github.com/ddan1l/tega-backend/usecases/project"
)

type UseCaseFactory interface {
	CreateAuthUseCase() auth_usecase.AuthUsecase
	CreateUserUseCase() project_usecase.ProjectUsecase
}

func (f *DefaultFactory) CreateAuthUseCase() auth_usecase.AuthUsecase {
	return auth_usecase.NewAuthUsecaseImpl(
		f.CreateUserRepository(),
		f.CreateTokenRepository(),
	)
}

func (f *DefaultFactory) CreateUserhUseCase() project_usecase.ProjectUsecase {
	return project_usecase.NewProjectUsecaseImpl(
		f.CreateProjectRepository(),
		f.CreateABAC(),
		f.CreateTxManager(),
	)
}
