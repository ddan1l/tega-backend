package auth_handler

import (
	"net/http"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"

	req "github.com/ddan1l/tega-backend/web/requests"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUsecase auth_usecase.AuthUsecase
}

func NewAuthHandler(authUsecase auth_usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var r req.RegisterUserRequest

	if !req.BindAndValidate(c, &r) {
		return
	}

	user := auth_dto.RegisterUserDto{
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
	}

	pair, err := h.authUsecase.RegisterUser(&user)

	if err != nil {
		res.ErrorResponse(c, err)
		return
	}

	h.SetAuthCookies(c, pair)

}

func (h *authHandler) SetAuthCookies(c *gin.Context, pair *auth_dto.TokensPairDto) {
	at := &http.Cookie{
		Name:     "AccessToken",
		Value:    pair.AccessToken,
		MaxAge:   60 * 15,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, at)

	rt := &http.Cookie{
		Name:     "RefreshToken",
		Value:    pair.RefreshToken,
		MaxAge:   3600 * 24,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, rt)
}
