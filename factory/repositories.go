package factory

import (
	token_repository "github.com/ddan1l/tega-backend/repositories/token"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
)

type RepositoryFactory interface {
	CreateUserRepository() user_repository.UserRepository
	CreateTokenRepository() token_repository.TokenRepository
}

func (f *DefaultFactory) CreateUserRepository() user_repository.UserRepository {
	return user_repository.NewUserPgRepository(f.db)
}

func (f *DefaultFactory) CreateTokenRepository() token_repository.TokenRepository {
	return token_repository.NewTokenPgRepository(f.db)
}
