package user_usercase

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
)

type UserUsecase interface {
	GetUserProjects(in *user_dto.FindByIdDto) (*user_dto.ProjectsDto, *errs.AppError)
	CreateProject(in *user_dto.CreateProjectDto) (*user_dto.ProjectDto, *errs.AppError)
}
