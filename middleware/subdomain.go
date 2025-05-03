package middleware

import (
	"github.com/ddan1l/tega-backend/ctx"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	errs "github.com/ddan1l/tega-backend/errors"
	user_usercase "github.com/ddan1l/tega-backend/usecases/user"
	res "github.com/ddan1l/tega-backend/web/responses"
	"github.com/gin-gonic/gin"
)

type subdomainMiddleware struct {
	userUseCase user_usercase.UserUsecase
}

func NewSubdomainMiddleware(userUseCase user_usercase.UserUsecase) Middleware {
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

			project, projectErr := m.userUseCase.CheckIsUserInProject(&project_dto.FindBySlugAndUserIdDto{
				UserID: user.ID,
				Slug:   subdomain,
			})

			if projectErr != nil {
				res.Error(c, projectErr)
				c.Abort()
				return
			}

			c.Set("Project", ctx.ProjectContext{
				ID:          project.ID,
				Name:        project.Name,
				Slug:        project.Slug,
				Description: project.Description,
			})
		}

		c.Next()
	}
}
