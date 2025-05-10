package res

import (
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
)

type ProjectsResponse struct {
	Projects []project_dto.ProjectDto `json:"projects"`
}

type ProjectResponse struct {
	Project project_dto.ProjectDto `json:"project"`
}

type ProjectPoliciesResponse struct {
	Policies abac_dto.PoliciesDto `json:"policies"`
}
