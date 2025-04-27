package user_handler

import (
	"github.com/ddan1l/tega-backend/ctx"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	user_usercase "github.com/ddan1l/tega-backend/usecases/user"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase user_usercase.UserUsecase
}

func NewUserHandler(userUsecase user_usercase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// Logout godoc
// @Summary User
// @Description AuthenticatedUser
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.UserResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Router /user [get]
func (h *userHandler) User(c *gin.Context) {
	var (
		user *ctx.UserContext
		err  error
	)

	user, err = ctx.GetUserFromContext(c)

	if err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
		return
	}

	r := &res.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}

	res.SuccessWithData(c, r)
}

// Logout godoc
// @Summary Projects
// @Description UserProjects
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.UserProjectsResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Response 400 {object} res.ErrorResponse{error=errs.BadRequestError}
// @Router /user/projects [get]
func (h *userHandler) UserProjects(c *gin.Context) {
	if user, err := ctx.GetUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {
		dto, err := h.userUsecase.GetUserProjects(&user_dto.FindByIdDto{ID: user.ID})

		if err != nil {
			res.Error(c, errs.BadRequest.WithMessage(err.Error()))
			return
		}

		r := &res.UserProjectsResponse{
			Projects: dto.Projects,
		}

		res.SuccessWithData(c, r)
	}
}
