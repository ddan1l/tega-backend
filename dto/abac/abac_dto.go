package abac_dto

import "github.com/ddan1l/tega-backend/models"

type (
	CreateDefaultPoliciesDto struct {
		ProjectID int
	}

	RoleDto struct {
		Role *models.Role
	}
)
