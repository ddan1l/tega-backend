package middleware

import (
	"net/http"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
	"github.com/ddan1l/tega-backend/utils"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	authUsecase auth_usecase.AuthUsecase
}

func NewAuthMiddleware(authUsecase auth_usecase.AuthUsecase) Middleware {
	return &authMiddleware{
		authUsecase: authUsecase,
	}
}

func (m *authMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		pair := auth_dto.TokensPairDto{
			AccessToken:  utils.SafeGetCookie(c, "AccessToken"),
			RefreshToken: utils.SafeGetCookie(c, "RefreshToken"),
		}

		r, err := m.authUsecase.Authenticate(&pair)

		if err != nil || r == nil || r.User == nil {
			m.clearCookies(c)

			res.Error(c, err)

			c.Abort()
			return
		}

		if r.AccessToken != "" {
			m.setAccessTokenCookie(c, r.AccessToken)
		}

		c.Set("User", user_dto.UserDto{
			ID:       r.User.ID,
			Email:    r.User.Email,
			FullName: r.User.FullName,
		})

		c.Next()
	}
}

func (m *authMiddleware) setAccessTokenCookie(c *gin.Context, token string) {
	at := &http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		MaxAge:   int(auth_usecase.AccessTokenMaxAge.Seconds()),
		Path:     "/",
		Domain:   ".tega.local",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, at)
}

func (m *authMiddleware) clearCookies(c *gin.Context) {
	at := &http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   ".tega.local",
		Secure:   false,
		HttpOnly: true,
	}

	rt := &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   ".tega.local",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, at)
	http.SetCookie(c.Writer, rt)
}
