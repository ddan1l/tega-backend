package auth_dto

import (
	"time"

	"github.com/ddan1l/tega-backend/models"
)

type (
	CreateTokenDto struct {
		UserID    int
		Token     string
		ExpiresAt time.Time
	}

	TokensPairDto struct {
		AccessToken  string
		RefreshToken string
	}

	TokensPairExpirationDto struct {
		AccessTokenExpiration  int
		RefreshTokenExpiration int
	}

	IssueTokenDto struct {
		UserID    int
		Secret    []byte
		ExpiresAt time.Time
	}

	ParseTokenDto struct {
		Token  string
		Secret []byte
	}

	AuthenticatedDto struct {
		User        *models.User
		AccessToken string
	}

	AuthenticatedUserDto struct {
		ID       int    `json:"id" example:"1"`
		FullName string `json:"fullname" example:"John"`
		Email    string `json:"email" example:"john@john.com"`
	}

	RegisterUserDto struct {
		FullName string
		Email    string
		Password string
	}

	LoginUserDto struct {
		Email    string
		Password string
	}
)
