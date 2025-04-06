package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model

	Id int

	RoleID int
	UserID int

	Role Role
	User User

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
