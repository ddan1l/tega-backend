package user_usercase

import (
	"github.com/ddan1l/tega-backend/abac"
	"github.com/ddan1l/tega-backend/database"
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	"github.com/ddan1l/tega-backend/models"
	project_repository "github.com/ddan1l/tega-backend/repositories/project"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
	"github.com/ddan1l/tega-backend/transaction"
)

type userUsecaseImpl struct {
	projectRepository project_repository.ProjectRepository
	userRepository    user_repository.UserRepository
	abac              abac.Engine
	txManager         transaction.TxManager
}

func NewUserUsecaseImpl(
	userRepository user_repository.UserRepository,
	projectRepository project_repository.ProjectRepository,
	abac abac.Engine,
	txManager transaction.TxManager,
) UserUsecase {
	return &userUsecaseImpl{
		userRepository:    userRepository,
		projectRepository: projectRepository,
		abac:              abac,
		txManager:         txManager,
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
	// Check is project exists
	if project, err := u.projectRepository.FindProjectsBySlug(&project_dto.FindBySlugDto{
		Slug: in.Project.Slug,
	}); err != nil || project != nil {
		if err != nil {
			return nil, errs.BadRequest.WithError(err)
		}

		if project != nil {
			return nil, errs.AlreadyExists.WithDetails(map[string]string{
				"slug": "Project with slug already exists.",
			})
		}
	}

	var (
		result  *project_dto.ProjectDto
		project *models.Project
		admin   *abac_dto.RoleDto
		err     error
	)

	// Run all in transaction
	txErr := u.txManager.CallWithTx(func(tx database.Database) *errs.AppError {

		// Create project
		project, err = u.projectRepository.WithTx(tx).CreateProject(in.Project)

		if err != nil {
			return errs.BadRequest.WithError(err)
		}

		// Create default policies
		admin, err = u.abac.WithTx(tx).CreateDefaultPolicies(&abac_dto.CreateDefaultPoliciesDto{
			ProjectID: project.ID,
		})

		if err != nil {
			return errs.BadRequest.WithError(err)
		}

		projectUser := &project_dto.ProjectUserDto{
			UserID:    in.User.ID,
			RoleID:    admin.Role.ID,
			ProjectID: project.ID,
		}

		// Create project user
		if _, err = u.projectRepository.WithTx(tx).CreateProjectUser(projectUser); err != nil {
			return errs.BadRequest.WithError(err)
		}

		result = &project_dto.ProjectDto{
			ID:          project.ID,
			Name:        project.Name,
			Slug:        project.Slug,
			Description: project.Description,
		}

		return nil
	})

	return result, txErr
}
