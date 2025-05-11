package project_usecase

import (
	"github.com/ddan1l/tega-backend/abac"
	"github.com/ddan1l/tega-backend/database"
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	errs "github.com/ddan1l/tega-backend/errors"
	"github.com/ddan1l/tega-backend/models"
	project_repository "github.com/ddan1l/tega-backend/repositories/project"
	"github.com/ddan1l/tega-backend/transaction"
)

type projectUsecaseImpl struct {
	projectRepository project_repository.ProjectRepository
	abac              abac.Engine
	txManager         transaction.TxManager
}

func NewProjectUsecaseImpl(
	projectRepository project_repository.ProjectRepository,
	abac abac.Engine,
	txManager transaction.TxManager,
) ProjectUsecase {
	return &projectUsecaseImpl{
		projectRepository: projectRepository,
		abac:              abac,
		txManager:         txManager,
	}
}

func (u *projectUsecaseImpl) GetProjectUser(in *project_dto.FindBySlugAndUserIdDto) (*project_dto.ProjectUserDto, *errs.AppError) {
	projectUser, err := u.projectRepository.FindProjectUser(in)

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	projectUserDto := &project_dto.ProjectUserDto{
		RoleID:    projectUser.RoleID,
		ProjectID: projectUser.ProjectID,
	}

	access := u.abac.CheckAccess(
		projectUserDto,
		models.ActionRead,
		models.ResourceProject,
	)

	if !access {
		return nil, errs.Forbidden.WithMessage("access denided")
	}

	return &project_dto.ProjectUserDto{
		ID:        projectUser.ID,
		UserID:    projectUser.UserID,
		RoleID:    projectUser.RoleID,
		ProjectID: projectUser.ProjectID,
		User: &user_dto.UserDto{
			ID:       projectUser.User.ID,
			FullName: projectUser.User.FullName,
			Email:    projectUser.User.Email,
		},
		Project: &project_dto.ProjectDto{
			ID:          projectUser.Project.ID,
			Slug:        projectUser.Project.Slug,
			Name:        projectUser.Project.Name,
			Description: projectUser.Project.Description,
		},
	}, nil
}

func (u *projectUsecaseImpl) GetProjectUsers(in *project_dto.ProjectUserDto) (*[]project_dto.ProjectUserDto, *errs.AppError) {
	projectUsers, err := u.projectRepository.FindProjectUsers(&project_dto.FindByIdDto{
		ID: in.ProjectID,
	})

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	access := u.abac.CheckAccess(
		in,
		models.ActionRead,
		models.ResourceUser,
	)

	if !access {
		return nil, errs.Forbidden.WithMessage("access denided")
	}

	users := make([]project_dto.ProjectUserDto, 0, len(*projectUsers))

	for _, projectUser := range *projectUsers {
		pu := project_dto.ProjectUserDto{
			ID:        projectUser.ID,
			UserID:    projectUser.UserID,
			RoleID:    projectUser.RoleID,
			ProjectID: projectUser.ProjectID,
			User: &user_dto.UserDto{
				ID:       projectUser.User.ID,
				FullName: projectUser.User.FullName,
				Email:    projectUser.User.Email,
			},
			Project: &project_dto.ProjectDto{
				ID:          projectUser.Project.ID,
				Slug:        projectUser.Project.Slug,
				Name:        projectUser.Project.Name,
				Description: projectUser.Project.Description,
			},
		}

		users = append(users, pu)
	}

	return &users, nil
}

func (u *projectUsecaseImpl) GetUserProjects(in *project_dto.FindByUserIdDto) (*project_dto.ProjectsDto, *errs.AppError) {
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

func (u *projectUsecaseImpl) CreateProject(in *project_dto.CreateProjectDto) (*project_dto.ProjectDto, *errs.AppError) {
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
			UserID:    in.UserID,
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

func (u *projectUsecaseImpl) GetProjectPolicies(in *project_dto.FindByIdDto) (*abac_dto.PoliciesDto, *errs.AppError) {
	res, err := u.abac.LoadProjectPolicies(&abac_dto.LoadProjectPoliciesDto{
		ProjectID: in.ID,
	})

	if err != nil {
		return nil, errs.BadRequest.WithError(err)
	}

	return res, nil
}
