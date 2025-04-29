package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ProjectUser struct {
	gorm.Model

	ID int `gorm:"primaryKey;autoIncrement"`

	UserID int
	User   User `gorm:"foreignKey:UserID"`

	ProjectID int
	Project   Project `gorm:"foreignKey:ProjectID"`

	RoleID int
	Role   Role `gorm:"foreignKey:RoleID"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
