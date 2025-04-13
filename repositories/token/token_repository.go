package token_repository

import (
	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	"github.com/ddan1l/tega-backend/models"
)

type TokenRepository interface {
	Create(in *auth_dto.CreateTokenDto) (*models.Token, error)
	Delete(t string) error
	FindByToken(t string) (*models.Token, error)
}
