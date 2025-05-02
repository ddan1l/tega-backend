package res

import (
	project_dto "github.com/ddan1l/tega-backend/dto/project"
)

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	FullName string `json:"fullName" example:"John"`
	Email    string `json:"email" example:"john@john.com"`
}

type UserProjectsResponse struct {
	Projects []project_dto.ProjectDto `json:"projects"`
}

type UserProjectResponse struct {
	Project project_dto.ProjectDto `json:"project"`
}
