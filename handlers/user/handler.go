package user_handler

import (
	"github.com/ddan1l/tega-backend/ctx"
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
