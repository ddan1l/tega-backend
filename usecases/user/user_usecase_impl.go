package user_usercase

import (
	project_dto "github.com/ddan1l/tega-backend/dto/project"
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

func (u *userUsecaseImpl) CheckIsUserInProject(in *project_dto.FindBySlugAndUserIdDto) (*project_dto.ProjectDto, *errs.AppError) {
	projects, err := u.projectRepository.FindProjectsByUserId(&user_dto.FindByIdDto{
		ID: in.UserID,
	})

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	for _, project := range *projects {
		if project.Slug == in.Slug {
			return &project_dto.ProjectDto{
				ID:          project.ID,
				Name:        project.Name,
				Slug:        project.Slug,
				Description: project.Description,
			}, nil
		}
	}

	return nil, errs.NotFound.WithMessage("Project not found.")
}

func (u *userUsecaseImpl) GetUserProjects(in *user_dto.FindByIdDto) (*project_dto.ProjectsDto, *errs.AppError) {
	projects, err := u.projectRepository.FindProjectsByUserId(in)

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	userProjects := make([]project_dto.ProjectDto, len(*projects))

	for i, project := range *projects {
		userProjects[i] = project_dto.ProjectDto{
			ID:          project.ID,
			Name:        project.Name,
			Slug:        project.Slug,
			Description: project.Description,
		}
	}

	return &project_dto.ProjectsDto{
		Projects: userProjects,
	}, nil
}

func (u *userUsecaseImpl) CreateProject(in *project_dto.CreateProjectDto) (*project_dto.ProjectDto, *errs.AppError) {
	project, err := u.projectRepository.FindProjectsBySlug(&project_dto.FindBySlugDto{
		Slug: in.Project.Slug,
	})

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	if project != nil {
		var details = make(map[string]string)

		details["slug"] = "Project with slug already exists."

		return nil, errs.AlreadyExists.WithDetails(
			details,
		)
	}

	if project, err := u.projectRepository.CreateProject(in.Project); err != nil {
		return nil, errs.BadRequest.WithError(err)
	} else {
		projectUser := &project_dto.ProjectUserDto{
			UserID:    in.User.ID,
			RoleID:    int(models.Owner),
			ProjectID: project.ID,
		}

		if _, err := u.projectRepository.CreateProjectUser(projectUser); err != nil {
			return nil, errs.BadRequest.WithError(err)
		} else {
			return &project_dto.ProjectDto{
				ID:          project.ID,
				Name:        project.Name,
				Slug:        project.Slug,
				Description: project.Description,
			}, nil
		}

	}

}
