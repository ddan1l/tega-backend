package middleware

import (
	"github.com/ddan1l/tega-backend/ctx"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	errs "github.com/ddan1l/tega-backend/errors"
	project_usecase "github.com/ddan1l/tega-backend/usecases/project"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type subdomainMiddleware struct {
	userUseCase project_usecase.ProjectUsecase
}

func NewSubdomainMiddleware(userUseCase project_usecase.ProjectUsecase) Middleware {
	return &subdomainMiddleware{
		userUseCase: userUseCase,
	}
}

func (m *subdomainMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		subdomain := c.Request.Header.Get("X-Subdomain")

		if subdomain != "" {
			user, userErr := ctx.GetUserFromContext(c)

			if userErr != nil {
				res.Error(c, errs.Forbidden.WithMessage(userErr.Error()))
				c.Abort()
				return
			}

			projectUser, projectErr := m.userUseCase.GetProjectUser(&project_dto.FindBySlugAndUserIdDto{
				UserID: user.ID,
				Slug:   subdomain,
			})

			if projectErr != nil {
				res.Error(c, projectErr)
				c.Abort()
				return
			}

			c.Set("ProjectUser", project_dto.ProjectUserDto{
				ID:        projectUser.ID,
				UserID:    projectUser.UserID,
				RoleID:    projectUser.RoleID,
				ProjectID: projectUser.ProjectID,
				Project: &project_dto.ProjectDto{
					ID:   projectUser.Project.ID,
					Slug: projectUser.Project.Slug,
				},
			})
		}

		c.Next()
	}
}
