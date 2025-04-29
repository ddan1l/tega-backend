package user_usercase

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	"github.com/ddan1l/tega-backend/models"
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
	projects, err := u.projectRepository.FindProjectsByUserId(in)

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	userProjects := make([]user_dto.ProjectDto, len(*projects))

	for i, project := range *projects {
		userProjects[i] = user_dto.ProjectDto{
			ID:          project.ID,
			Name:        project.Name,
			Slug:        project.Slug,
			Description: project.Description,
		}
	}

	return &user_dto.ProjectsDto{
		Projects: userProjects,
	}, nil
}

func (u *userUsecaseImpl) CreateProject(in *user_dto.CreateProjectDto) (*user_dto.ProjectDto, *errs.AppError) {
	if project, err := u.projectRepository.CreateProject(in.Project); err != nil {
		return nil, errs.BadRequest.WithError(err)
	} else {
		projectUser := &user_dto.ProjectUserDto{
			UserID:    in.User.ID,
			RoleID:    int(models.Owner),
			ProjectID: project.ID,
		}

		if _, err := u.projectRepository.CreateProjectUser(projectUser); err != nil {
			return nil, errs.BadRequest.WithError(err)
		} else {
			return &user_dto.ProjectDto{
				Name:        project.Name,
				Slug:        project.Slug,
				Description: project.Description,
			}, nil
		}

	}

}
