package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string
	Slug string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
