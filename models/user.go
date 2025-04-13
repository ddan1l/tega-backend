package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id           int `gorm:"primaryKey;autoIncrement"`
	FullName     string
	Email        string
	PasswordHash string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
