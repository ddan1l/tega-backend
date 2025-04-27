package user_usercase

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	project_repository "github.com/ddan1l/tega-backend/repositories/project"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
)

type userUsecaseImpl struct {
	projectRepository project_repository.ProjectRepository
	userRepository    user_repository.UserRepository
}

func NewUserUsecaseImpl(
	userRepository user_repository.UserRepository,
	projectRepository project_repository.ProjectRepository,
) UserUsecase {
	return &userUsecaseImpl{
		userRepository:    userRepository,
		projectRepository: projectRepository,
	}
}

func (u *userUsecaseImpl) GetUserProjects(in *user_dto.FindByIdDto) (*user_dto.ProjectsDto, *errs.AppError) {
	projects, err := u.projectRepository.FindByUserId(in)

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	userProjects := make([]user_dto.ProjectDto, len(*projects))

	for i, project := range *projects {
		userProjects[i] = user_dto.ProjectDto{
			ID:   project.ID,
			Name: project.Name,
			Slug: project.Slug,
		}
	}

	return &user_dto.ProjectsDto{
		Projects: userProjects,
	}, nil
}
