package user_handler

import (
	"github.com/ddan1l/tega-backend/ctx"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	project_usecase "github.com/ddan1l/tega-backend/usecases/project"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	projectUsecase project_usecase.ProjectUsecase
}

func NewUserHandler(projectUsecase project_usecase.ProjectUsecase) UserHandler {
	return &userHandler{
		projectUsecase: projectUsecase,
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
		user *user_dto.UserDto
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
// @Summary User app
// @Description User app
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.UserAppResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Router /user/app [get]
func (h *userHandler) UserApp(c *gin.Context) {
	if projectUser, err := ctx.GetProjectUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
		return
	} else {
		userProjects, err := h.projectUsecase.GetUserProjects(&project_dto.FindByUserIdDto{UserID: projectUser.UserID})

		if err != nil {
			res.Error(c, errs.Forbidden.WithMessage(err.Error()))
			return
		}

		r := &res.UserAppResponse{
			ProjectUser: *projectUser,
			Projects:    userProjects.Projects,
		}

		res.SuccessWithData(c, r)
	}

}
