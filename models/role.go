package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ID int

const (
	Owner  ID = 1
	Member ID = 2
	Viewer ID = 3
)

type Role struct {
	gorm.Model

	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
