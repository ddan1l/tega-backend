package abac

import (
	"github.com/ddan1l/tega-backend/database"
	abac_dto "github.com/ddan1l/tega-backend/dto/abac"
	"github.com/ddan1l/tega-backend/models"
	"gorm.io/gorm"
)

type engine struct {
	db database.Database
}

func NewEngine(db database.Database) Engine {
	return &engine{db: db}
}

func (e *engine) WithTx(db database.Database) Engine {
	return &engine{db: db}
}

func (e *engine) CreateDefaultPolicies(in *abac_dto.CreateDefaultPoliciesDto) (*abac_dto.RoleDto, error) {
	// Создаем роли
	roles := []models.Role{
		{Slug: "admin", ProjectID: in.ProjectID},
		{Slug: "developer", ProjectID: in.ProjectID},
		{Slug: "viewer", ProjectID: in.ProjectID},
	}

	if err := e.db.GetDb().Create(&roles).Error; err != nil {
		return nil, err
	}

	adminRole, developerRole, viewerRole := roles[0], roles[1], roles[2]

	// Создаем политики с использованием ENUM типов
	policies := []models.Policy{
		{
			Slug:      "admin-full-access",
			RoleID:    adminRole.ID,
			ProjectID: in.ProjectID,
			Actions: []models.PolicyAction{
				{Action: models.ActionCreate},
				{Action: models.ActionRead},
				{Action: models.ActionUpdate},
				{Action: models.ActionDelete},
			},
			Resources: []models.PolicyResource{
				{Resource: models.ResourceProject},
				{Resource: models.ResourceUser},
				{Resource: models.ResourceTask},
			},
		},
		{
			Slug:      "developer-task-access",
			RoleID:    developerRole.ID,
			ProjectID: in.ProjectID,
			Actions: []models.PolicyAction{
				{Action: models.ActionCreate},
				{Action: models.ActionRead},
				{Action: models.ActionUpdate},
				{Action: models.ActionDelete},
			},
			Resources: []models.PolicyResource{
				{Resource: models.ResourceTask},
			},
		},
		{
			Slug:      "developer-view-project",
			RoleID:    developerRole.ID,
			ProjectID: in.ProjectID,
			Actions: []models.PolicyAction{
				{Action: models.ActionRead},
			},
			Resources: []models.PolicyResource{
				{Resource: models.ResourceTask},
			},
		},
		{
			Slug:      "viewer-read-only",
			RoleID:    viewerRole.ID,
			ProjectID: in.ProjectID,
			Actions: []models.PolicyAction{
				{Action: models.ActionRead},
			},
			Resources: []models.PolicyResource{
				{Resource: models.ResourceProject},
				{Resource: models.ResourceTask},
			},
		},
	}

	err := e.db.GetDb().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&policies).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &abac_dto.RoleDto{
		Role: &adminRole,
	}, nil
}

//
//
//
//
//
//
//
//
//
//
//
//
////
//
////
//
////
//
//////
//
//
//
////
//
//

//

// func checkAccess(policies []models.Policy, user models.ProjectUser, action string, resource string) bool {
// 	for _, policy := range policies {
// 		// Проверяем роль
// 		if policy.Role != user.Role {
// 			continue
// 		}

// 		// Проверяем действие
// 		actionAllowed := false
// 		for _, a := range policy.Actions {
// 			if a == action {
// 				actionAllowed = true
// 				break
// 			}
// 		}
// 		if !actionAllowed {
// 			continue
// 		}

// 		// Проверяем ресурс
// 		resourceAllowed := false
// 		for _, r := range policy.Resources {
// 			if r == resource {
// 				resourceAllowed = true
// 				break
// 			}
// 		}
// 		if !resourceAllowed {
// 			continue
// 		}

// 		// Если все проверки пройдены
// 		return policy.Effect == "allow"
// 	}

// 	return false
// }
