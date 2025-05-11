package project_handler

import (
	"github.com/ddan1l/tega-backend/ctx"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	errs "github.com/ddan1l/tega-backend/errors"
	project_usecase "github.com/ddan1l/tega-backend/usecases/project"
	req "github.com/ddan1l/tega-backend/web/requests"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	projectUsecase project_usecase.ProjectUsecase
}

func NewProjectHandler(projectUsecase project_usecase.ProjectUsecase) UserHandler {
	return &userHandler{
		projectUsecase: projectUsecase,
	}
}

// Logout godoc
// @Summary Projects
// @Description UserProjects
// @Tags project
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.ProjectsResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Response 400 {object} res.ErrorResponse{error=errs.BadRequestError}
// @Router /projects [get]
func (h *userHandler) UserProjects(c *gin.Context) {
	if user, err := ctx.GetUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {
		dto, err := h.projectUsecase.GetUserProjects(&project_dto.FindByUserIdDto{UserID: user.ID})

		if err != nil {
			res.Error(c, errs.BadRequest.WithMessage(err.Error()))
			return
		}

		r := &res.ProjectsResponse{
			Projects: dto.Projects,
		}

		res.SuccessWithData(c, r)
	}
}

// Logout godoc
// @Summary Create Project
// @Description Create Project
// @Tags project
// @Accept json
// @Produce json
// @Param request body req.CreateProjectRequest true "Create project request"
// @Success 200 {object} res.SuccessWithDataResponse{data=res.ProjectResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Response 400 {object} res.ErrorResponse{error=errs.BadRequestError}
// @Response 422 {object} res.ErrorResponse{error=errs.ValidationFailedError}
// @Router /project [post]
func (h *userHandler) CreateProject(c *gin.Context) {
	var r req.CreateProjectRequest

	if !req.BindAndValidate(c, &r) {
		return
	}

	if user, err := ctx.GetUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {

		createProjectDto := &project_dto.CreateProjectDto{
			Project: &project_dto.ProjectDto{
				Name:        r.Name,
				Slug:        r.Slug,
				Description: r.Description,
			},
			UserID: user.ID,
		}

		dto, err := h.projectUsecase.CreateProject(createProjectDto)

		if err != nil {
			res.Error(c, err)
			return
		}

		r := &res.ProjectResponse{
			Project: *dto,
		}

		res.SuccessWithData(c, r)
	}
}

func (h *userHandler) ProjectsPolicies(c *gin.Context) {
	if projectUser, err := ctx.GetProjectUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {
		dto, err := h.projectUsecase.GetProjectPolicies(&project_dto.FindByIdDto{
			ID: projectUser.ProjectID,
		})

		if err != nil {
			res.Error(c, err)
			return
		}

		r := &res.ProjectPoliciesResponse{
			Policies: dto.Policies,
		}

		res.SuccessWithData(c, r)
	}
}

// Logout godoc
// @Summary Project user
// @Description Project user
// @Tags project
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.ProjectUserResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Response 400 {object} res.ErrorResponse{error=errs.BadRequestError}
// @Router /project/user [get]
func (h *userHandler) ProjectUser(c *gin.Context) {
	if user, err := ctx.GetProjectUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {

		r := &res.ProjectUserResponse{
			User: *user,
		}

		res.SuccessWithData(c, r)
	}
}

// Logout godoc
// @Summary Project users
// @Description Project users
// @Tags project
// @Accept json
// @Produce json
// @Success 200 {object} res.SuccessWithDataResponse{data=res.ProjectUsersResponse}
// @Response 403 {object} res.ErrorResponse{error=errs.ForbiddenError}
// @Response 400 {object} res.ErrorResponse{error=errs.BadRequestError}
// @Router /project/users [get]
func (h *userHandler) ProjectUsers(c *gin.Context) {
	if user, err := ctx.GetProjectUserFromContext(c); err != nil {
		res.Error(c, errs.Forbidden.WithMessage(err.Error()))
	} else {
		dto, err := h.projectUsecase.GetProjectUsers(user)

		if err != nil {
			res.Error(c, err)
			return
		}

		r := &res.ProjectUsersResponse{
			Users: *dto,
		}

		res.SuccessWithData(c, r)
	}
}
