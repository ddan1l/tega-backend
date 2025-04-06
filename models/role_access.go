package models

import (
	"time"

	"gorm.io/gorm"
)

type RoleAccess struct {
	gorm.Model

	Id int

	RoleID   int
	AccessId int

	Role   Role
	Access Access

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
