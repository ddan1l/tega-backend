package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Name string

const (
	Owner  Name = "owner"
	Member Name = "member"
	Viewer Name = "viewer"
)

type Role struct {
	gorm.Model

	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
