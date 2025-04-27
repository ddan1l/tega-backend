package res

import (
	user_dto "github.com/ddan1l/tega-backend/dto/user"
)

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	FullName string `json:"fullName" example:"John"`
	Email    string `json:"email" example:"john@john.com"`
}

type UserProjectsResponse struct {
	Projects []user_dto.ProjectDto `json:"projects"`
}
