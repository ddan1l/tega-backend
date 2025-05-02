package user_usercase

import (
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
)

type UserUsecase interface {
	GetUserProjects(in *user_dto.FindByIdDto) (*project_dto.ProjectsDto, *errs.AppError)
	CreateProject(in *project_dto.CreateProjectDto) (*project_dto.ProjectDto, *errs.AppError)
}
