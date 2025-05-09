package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ID int

type Role struct {
	gorm.Model

	ID    int `gorm:"primaryKey;autoIncrement"`
	Title string
	Slug  string

	ProjectID int
	Project   Project `gorm:"foreignKey:ProjectID"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
