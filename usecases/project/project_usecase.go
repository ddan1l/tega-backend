package project_usecase

import (
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	errs "github.com/ddan1l/tega-backend/errors"
)

type ProjectUsecase interface {
	CheckIsUserInProject(in *project_dto.FindBySlugAndUserIdDto) (*project_dto.ProjectDto, *errs.AppError)
	GetUserProjects(in *project_dto.FindByUserIdDto) (*project_dto.ProjectsDto, *errs.AppError)
	CreateProject(in *project_dto.CreateProjectDto) (*project_dto.ProjectDto, *errs.AppError)
	GetProjectPolicies(in *project_dto.FindByIdDto) (*abac_dto.PoliciesDto, *errs.AppError)
}
