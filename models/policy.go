package models

import (
	"gorm.io/gorm"
)

type ActionType string
type ResourceType string
type ConditionOperator string

const (
	ActionCreate ActionType = "create"
	ActionRead   ActionType = "read"
	ActionUpdate ActionType = "update"
	ActionDelete ActionType = "delete"

	ResourceProject ResourceType = "project"
	ResourceTask    ResourceType = "task"
	ResourceUser    ResourceType = "user"

	OperatorEquals     ConditionOperator = "eq"
	OperatorNotEquals  ConditionOperator = "neq"
	OperatorContains   ConditionOperator = "contains"
	OperatorStartsWith ConditionOperator = "startsWith"
)

type Policy struct {
	gorm.Model

	Slug   string `gorm:"size:100;not null"`
	Effect string `gorm:"size:20;not null;check:effect IN ('allow','deny');default:allow"`

	RoleID int  `gorm:"not null"`
	Role   Role `gorm:"foreignKey:RoleID"`

	ProjectID int     `gorm:"not null"`
	Project   Project `gorm:"foreignKey:ProjectID"`

	Actions    []PolicyAction    `gorm:"foreignKey:PolicyID;constraint:OnDelete:CASCADE"`
	Resources  []PolicyResource  `gorm:"foreignKey:PolicyID;constraint:OnDelete:CASCADE"`
	Conditions []PolicyCondition `gorm:"foreignKey:PolicyID;constraint:OnDelete:CASCADE"`
}

type PolicyAction struct {
	gorm.Model

	PolicyID uint       `gorm:"index;not null"`
	Action   ActionType `gorm:"size:50;not null"`
}

type PolicyResource struct {
	gorm.Model

	PolicyID uint         `gorm:"index;not null"`
	Resource ResourceType `gorm:"size:50;not null"`
}

type PolicyCondition struct {
	gorm.Model

	PolicyID uint `gorm:"index;not null"`

	Field    string            `gorm:"size:100;not null"`
	Operator ConditionOperator `gorm:"size:20;not null"`
	Value    string            `gorm:"size:255;not null"`
}
