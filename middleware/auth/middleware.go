package auth_middleware

import (
	"net/http"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
	"github.com/ddan1l/tega-backend/utils"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	authUsecase auth_usecase.AuthUsecase
}

func NewAuthMiddleware(authUsecase auth_usecase.AuthUsecase) AuthMiddleware {
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

		authenticatedDto, err := m.authUsecase.Authenticate(&pair)

		if err != nil || authenticatedDto == nil || authenticatedDto.User == nil {
			m.clearCookies(c)

			res.Error(c, err)

			c.Abort()
			return
		}

		if authenticatedDto.AccessToken != "" {
			m.setAccessTokenCookie(c, authenticatedDto.AccessToken)
		}

		c.Set("user", authenticatedDto.User)
		c.Next()
	}
}

func (m *authMiddleware) setAccessTokenCookie(c *gin.Context, token string) {
	at := &http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		MaxAge:   int(auth_usecase.AccessTokenMaxAge.Seconds()),
		Path:     "/",
		Domain:   "",
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
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}

	rt := &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, at)
	http.SetCookie(c.Writer, rt)
}
