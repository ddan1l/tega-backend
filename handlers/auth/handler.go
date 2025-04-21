package auth_handler

import (
	"net/http"

	auth_dto "github.com/ddan1l/tega-backend/dto/auth"
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
	"github.com/ddan1l/tega-backend/utils"

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

// Register godoc
// @Summary User register
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.RegisterUserRequest true "Register credentials"
// @Success 200 {object} res.SuccessResponse
// @Response 404 {object} res.ErrorResponse{error=errs.AppError}
// @Response 409 {object} res.ErrorResponse{error=errs.AlreadyExistsError}
// @Response 422 {object} res.ErrorResponse{error=errs.ValidationFailedError}
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var r req.RegisterUserRequest

	if !req.BindAndValidate(c, &r) {
		return
	}

	dto := auth_dto.RegisterUserDto{
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
	}

	if pair, err := h.authUsecase.RegisterUser(&dto); err != nil {
		res.Error(c, err)
	} else {
		exp := &auth_dto.TokensPairExpirationDto{
			AccessTokenExpiration:  int(auth_usecase.AccessTokenMaxAge.Seconds()),
			RefreshTokenExpiration: int(auth_usecase.RefreshTokenMaxAge.Seconds()),
		}

		h.setAuthCookies(c, pair, exp)

		res.Succes(c)
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.LoginUserRequest true "Login credentials"
// @Success 200 {object} res.SuccessResponse
// @Response 404 {object} res.ErrorResponse{error=errs.UserNotFoundError}
// @Response 401 {object} res.ErrorResponse{error=errs.IncorrectPasswordError}
// @Response 422 {object} res.ErrorResponse{error=errs.ValidationFailedError}
// @Router /auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var r req.LoginUserRequest

	if !req.BindAndValidate(c, &r) {
		return
	}

	dto := auth_dto.LoginUserDto{
		Email:    r.Email,
		Password: r.Password,
	}

	if pair, err := h.authUsecase.LoginUser(&dto); err != nil {
		res.Error(c, err)
	} else {
		exp := &auth_dto.TokensPairExpirationDto{
			AccessTokenExpiration:  int(auth_usecase.AccessTokenMaxAge.Seconds()),
			RefreshTokenExpiration: int(auth_usecase.RefreshTokenMaxAge.Seconds()),
		}

		h.setAuthCookies(c, pair, exp)

		res.Succes(c)
	}
}

// Logout godoc
// @Summary User logout
// @Description Logout user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessResponse
// @Response 400 {object} res.ErrorResponse{error=errs.AuthError}
// @Router /auth/logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	refreshToken := utils.SafeGetCookie(c, "RefreshToken")

	pair := &auth_dto.TokensPairDto{
		AccessToken:  "",
		RefreshToken: "",
	}

	exp := &auth_dto.TokensPairExpirationDto{
		AccessTokenExpiration:  -1,
		RefreshTokenExpiration: -1,
	}

	h.setAuthCookies(c, pair, exp)

	if refreshToken != "" {
		if err := h.authUsecase.DeleteToken(refreshToken); err != nil {
			res.Error(c, err)
			return
		}
	}

	res.Succes(c)
}

// Logout godoc
// @Summary User
// @Description AuthenticatedUser
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessResponse
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Router /auth/user [get]
func (h *authHandler) User(c *gin.Context) {
	user, _ := c.Get("User")

	// if !exists {
	// 	res.Error(c, errs.Forbidden)
	// 	return
	// }

	res.SuccessWithData(c, user)
}

func (h *authHandler) setAuthCookies(c *gin.Context, pair *auth_dto.TokensPairDto, exp *auth_dto.TokensPairExpirationDto) {

	at := &http.Cookie{
		Name:     "AccessToken",
		Value:    pair.AccessToken,
		MaxAge:   exp.AccessTokenExpiration,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}

	rt := &http.Cookie{
		Name:     "RefreshToken",
		Value:    pair.RefreshToken,
		MaxAge:   exp.RefreshTokenExpiration,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, at)
	http.SetCookie(c.Writer, rt)
}
