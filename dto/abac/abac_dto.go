package abac_dto

import "github.com/ddan1l/tega-backend/models"

type (
	CreateDefaultPoliciesDto struct {
		ProjectID int
	}

	LoadProjectPoliciesDto struct {
		ProjectID int
	}

	PoliciesDto struct {
		Policies *[]models.Policy
	}

	RoleDto struct {
		Role *models.Role
	}
)
