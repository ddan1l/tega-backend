package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           int `gorm:"primaryKey;autoIncrement"`
	FullName     string
	Email        string `gorm:"unique"`
	PasswordHash string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
