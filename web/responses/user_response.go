package res

import project_dto "github.com/ddan1l/tega-backend/dto/project"

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	FullName string `json:"fullName" example:"John"`
	Email    string `json:"email" example:"john@john.com"`
}

type UserAppResponse struct {
	ProjectUser project_dto.ProjectUserDto `json:"projectUser"`
	Projects    []project_dto.ProjectDto   `json:"projects"`
}
