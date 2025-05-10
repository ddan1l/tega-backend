package res

import (
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"github.com/ddan1l/tega-backend/models"
)

type ProjectsResponse struct {
	Projects []project_dto.ProjectDto `json:"projects"`
}

type ProjectResponse struct {
	Project project_dto.ProjectDto `json:"project"`
}

type ProjectPoliciesResponse struct {
	Policies *[]models.Policy `json:"policies"`
}
