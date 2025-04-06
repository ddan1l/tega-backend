package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id           int `gorm:"index"`
	FullName     string
	Email        string
	PasswordHash string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
