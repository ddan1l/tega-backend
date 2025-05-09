package abac

import (
	"github.com/ddan1l/tega-backend/database"
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
)

type Engine interface {
	CreateDefaultPolicies(in *abac_dto.CreateDefaultPoliciesDto) (*abac_dto.RoleDto, error)
	WithTx(db database.Database) Engine
}
