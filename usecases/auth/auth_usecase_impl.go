package auth_usecase

import (
	"errors"
	"time"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	token_repository "github.com/ddan1l/tega-backend/repositories/token"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
	"github.com/ddan1l/tega-backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("access_secret_key2")
var refreshSecret = []byte("refresh_secret_key")

type JwtClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

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

// RegisterUser creates a new user account and returns JWT tokens.
// Performs email availability check and password hashing.
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
				UserId:    user.Id,
				Token:     pair.RefreshToken,
				ExpiresAt: time.Now().Add(RefreshTokenMaxAge),
			})

			if err != nil {
				return nil, errs.Auth.WithError(err)
			}

			return pair, nil
		}
	}
}

// LoginUser authenticates a user by email/password and returns new JWT tokens.
func (u *authUsecaseImpl) LoginUser(in *auth_dto.LoginUserDto) (*auth_dto.TokensPairDto, *errs.AppError) {
	user, err := u.userRepository.FindByEmail(&user_dto.FindByEmailDto{
		Email: in.Email,
	})

	if err != nil {
		return nil, errs.Auth.WithError(err)
	}

	if user == nil {
		return nil, errs.UserNotFound
	}

	if ok := utils.CheckPasswordHash(in.Password, user.PasswordHash); !ok {
		return nil, errs.IncorrectPassword
	} else {
		if pair, err := u.issuePair(user.Id); err != nil {
			return nil, errs.Auth.WithError(err)
		} else {
			_, err := u.tokenRepository.Create(&auth_dto.CreateTokenDto{
				UserId:    user.Id,
				Token:     pair.RefreshToken,
				ExpiresAt: time.Now().Add(RefreshTokenMaxAge),
			})

			if err != nil {
				return nil, errs.Auth.WithError(err)
			}

			return pair, nil
		}
	}

}

// CheckUserExists verifies if a user with the given email already exists.
// Returns:
//   - nil if email is available
//   - AlreadyExists if user found
//   - Auth-wrapped error for repository failures
func (u *authUsecaseImpl) CheckUserExists(e string) *errs.AppError {
	user, err := u.userRepository.FindByEmail(&user_dto.FindByEmailDto{
		Email: e,
	})

	if err != nil {
		return errs.Auth.WithError(err)
	}

	if user != nil {
		return errs.AlreadyExists
	}

	return nil
}

// issuePair generates a new pair of JWT tokens (access and refresh) for the given user ID.
// Returns:
//   - TokensPairDto containing both tokens on success
//   - Auth-wrapped error if either token generation fails
func (u *authUsecaseImpl) issuePair(id int) (*auth_dto.TokensPairDto, *errs.AppError) {

	accessToken, err := u.issueToken(&auth_dto.IssueTokenDto{
		UserId:    id,
		ExpiresAt: time.Now().Add(AccessTokenMaxAge),
		Secret:    accessSecret,
	})

	if err != nil {
		return nil, errs.Auth.WithError(err)
	}

	refreshToken, err := u.issueToken(&auth_dto.IssueTokenDto{
		UserId:    id,
		ExpiresAt: time.Now().Add(RefreshTokenMaxAge),
		Secret:    refreshSecret,
	})

	if err != nil {
		return nil, errs.Auth.WithError(err)
	}

	return &auth_dto.TokensPairDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// issueToken generates a JWT token for the specified user ID with given secret key.
// The token contains:
//   - user_id claim (provided userId)
//   - exp claim (15 minutes from issuance)
//
// Returns:
//   - Signed token string on success
//   - Auth-wrapped error if signing fails
func (u *authUsecaseImpl) issueToken(in *auth_dto.IssueTokenDto) (string, *errs.AppError) {

	claims := JwtClaims{
		UserId: in.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(in.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(in.Secret)

	if err != nil {
		return "", errs.Auth.WithError(err)
	} else {
		return tokenString, nil
	}
}

// parseToken validates and parses a JWT token using the provided secret.
// Returns:
//   - JwtClaims if token is valid
//   - AppError with:
//   - TokenExpired if token expired
//   - Forbidden for any other validation errors (invalid signature, malformed token, etc.)
func (u *authUsecaseImpl) parseToken(in *auth_dto.ParseTokenDto) (*JwtClaims, *errs.AppError) {
	token, err := jwt.ParseWithClaims(in.Token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return in.Secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errs.TokenExpired
		}
		return nil, errs.Forbidden.WithError(err)
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errs.Forbidden.WithError(err)
}

// Authenticate handles user authentication using either access or refresh token
// Returns authenticated user data with new access token if refresh was used
func (u *authUsecaseImpl) Authenticate(in *auth_dto.TokensPairDto) (*auth_dto.AuthenticatedDto, *errs.AppError) {
	// Initialize validation flags and output variables
	var (
		isEmptyAccess  = in.AccessToken == ""  // Flag for empty access token
		isEmptyRefresh = in.RefreshToken == "" // Flag for empty refresh token
		isEmptyPair    = isEmptyAccess && isEmptyRefresh
		newAccessToken = "" // Will contain new access token if generated
		userId         = 0  // Will store authenticated user ID
	)

	// Reject requests with no valid tokens
	if isEmptyPair || isEmptyRefresh {
		return nil, errs.Forbidden
	}

	// Case 1: Only refresh token provided (access token expired/missing)
	if isEmptyAccess {
		// Validate refresh token structure and signature
		_, errRefreshToken := u.parseToken(&auth_dto.ParseTokenDto{
			Token:  in.RefreshToken,
			Secret: refreshSecret,
		})

		if errRefreshToken != nil {
			return nil, errRefreshToken
		}

		// Check if refresh token exists in database (not revoked)
		refreshToken, errFindToken := u.tokenRepository.FindByToken(in.RefreshToken)
		if errFindToken != nil {
			return nil, errs.Forbidden.WithError(errFindToken)
		}

		// Prepare data for new access token generation
		newAccessTokenDto := &auth_dto.IssueTokenDto{
			UserId:    refreshToken.UserId, // Use ID from refresh token
			ExpiresAt: time.Now().Add(AccessTokenMaxAge),
			Secret:    accessSecret,
		}

		// Generate new access token
		if accessToken, err := u.issueToken(newAccessTokenDto); err != nil {
			return nil, errs.Forbidden.WithError(err)
		} else {
			newAccessToken = accessToken
			userId = refreshToken.UserId
		}

	} else {
		// Case 2: Valid access token provided
		claims, errAccessToken := u.parseToken(&auth_dto.ParseTokenDto{
			Token:  in.AccessToken,
			Secret: accessSecret,
		})

		if errAccessToken != nil {
			return nil, errAccessToken
		}

		userId = claims.UserId // Extract user ID from access token claims
	}

	// Fetch full user data by ID obtained from tokens
	user, errFindUser := u.userRepository.FindById(&user_dto.FindByIdDto{
		Id: userId,
	})

	if errFindUser != nil {
		return nil, errs.Forbidden.WithError(errFindUser)
	}

	// Prepare successful response
	result := auth_dto.AuthenticatedDto{
		User:        user,           // User profile data
		AccessToken: newAccessToken, // New access token (empty if original was valid)
	}

	return &result, nil
}

// Delete token from database
func (u *authUsecaseImpl) DeleteToken(t string) *errs.AppError {
	err := u.tokenRepository.Delete(t)

	if err != nil {
		return errs.Auth.WithError(err)
	}

	return nil
}
