package abac

import (
	"github.com/ddan1l/tega-backend/database"
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"github.com/ddan1l/tega-backend/models"
)

type Engine interface {
	CreateDefaultPolicies(in *abac_dto.CreateDefaultPoliciesDto) (*abac_dto.RoleDto, error)
	LoadProjectPolicies(in *abac_dto.LoadProjectPoliciesDto) (*abac_dto.PoliciesDto, error)
	CheckAccess(user *project_dto.ProjectUserDto, action models.ActionType, resource models.ResourceType)
	WithTx(db database.Database) Engine
}
