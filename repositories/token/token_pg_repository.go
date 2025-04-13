package token_repository

import (
	"github.com/ddan1l/tega-backend/database"
	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	"github.com/ddan1l/tega-backend/models"
)

type tokenPgRepository struct {
	db database.Database
}

func NewTokenPgRepository(db database.Database) TokenRepository {
	return &tokenPgRepository{db: db}
}

func (r *tokenPgRepository) Create(in *auth_dto.CreateTokenDto) (*models.Token, error) {
	token := &models.Token{
		Token:     in.Token,
		UserId:    in.UserId,
		ExpiresAt: in.ExpiresAt,
	}

	result := r.db.GetDb().Create(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}

func (r *tokenPgRepository) FindByToken(t string) (*models.Token, error) {
	var token models.Token

	result := r.db.GetDb().Where(models.Token{
		Token: t,
	}).First(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	return &token, nil
}

func (r *tokenPgRepository) Delete(t string) error {
	var token models.Token

	result := r.db.GetDb().Where(models.Token{
		Token: t,
	}).Delete(&token)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
