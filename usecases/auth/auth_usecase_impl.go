package auth_usecase

import (
	"time"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	token_repository "github.com/ddan1l/tega-backend/repositories/token"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
	"github.com/ddan1l/tega-backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("access_secret_key")
var refreshSecret = []byte("refresh_secret_key")

type authUsecaseImpl struct {
	userRepository  user_repository.UserRepository
	tokenRepository token_repository.TokenRepository
}

func NewAuthUsecaseImpl(
	userRepository user_repository.UserRepository,
	tokenRepository token_repository.TokenRepository,
) AuthUsecase {
	return &authUsecaseImpl{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
	}
}

// RegisterUser registers a new user by first checking if the user already exists,
// then hashing the user's password, creating the user in the repository, and issuing a token pair.
// It returns a pair of tokens (access and refresh tokens) if successful, or an error if any step fails.
func (u *authUsecaseImpl) RegisterUser(in *auth_dto.RegisterUserDto) (*auth_dto.TokensPairDto, *errs.AppError) {

	if err := u.CheckUserExists(in.Email); err != nil {
		return nil, err
	}

	if hash, err := utils.HashPassword(in.Password); err != nil {
		return nil, errs.Auth.WithError(err)
	} else {
		user, err := u.userRepository.Create(&user_dto.CreateUserDto{
			Email:        in.Email,
			FullName:     in.FullName,
			PasswordHash: hash,
		})

		if err != nil {
			return nil, errs.Auth.WithError(err)
		}

		if pair, err := u.issuePair(user.Id); err != nil {
			return nil, errs.Auth.WithError(err)
		} else {
			_, err := u.tokenRepository.Create(&auth_dto.CreateTokenDto{
				UserId: user.Id,
				Token:  pair.RefreshToken,
			})

			if err != nil {
				return nil, errs.Auth.WithError(err)
			}

			return pair, nil
		}
	}

}

// CheckUserExists checks whether a user already exists in the system based on their email address.
// It queries the user repository to find a user by the provided email.
// If the user exists, it returns an error indicating that the user already exists.
// If the user does not exist, it returns nil, indicating that the email is available for registration.
func (u *authUsecaseImpl) CheckUserExists(email string) *errs.AppError {
	user, err := u.userRepository.FindByEmail(&user_dto.FindByEmailDto{
		Email: email,
	})

	if err != nil {
		return errs.Auth.WithError(err)
	}

	if user != nil {
		return errs.AlreadyExists
	}

	return nil
}

// issuePair generates a pair of JWT tokens: an access token and a refresh token for the given user ID.
// It calls issueToken to generate each token using different secrets (accessSecret and refreshSecret).
// If either token generation fails, it returns an authentication error.
// On success, it returns a TokensPairDto containing both tokens.
func (u *authUsecaseImpl) issuePair(userId int) (*auth_dto.TokensPairDto, *errs.AppError) {

	accessToken, err := u.issueToken(userId, accessSecret)

	if err != nil {
		return nil, errs.Auth.WithError(err)
	}

	refreshToken, err := u.issueToken(userId, refreshSecret)

	if err != nil {
		return nil, errs.Auth.WithError(err)
	}

	return &auth_dto.TokensPairDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// issueToken generates a new JWT token with the given user ID and secret.
// The token is signed using the HS256 algorithm and includes an expiration time of 15 minutes.
// If token creation or signing fails, it returns an authentication error.
// On success, it returns the generated token as a string.
func (u *authUsecaseImpl) issueToken(userId int, secret []byte) (string, *errs.AppError) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	t, err := token.SignedString(secret)

	if err != nil {
		return "", errs.Auth.WithError(err)
	} else {
		return t, nil
	}
}
