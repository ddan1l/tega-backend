package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id           int
	Name         string
	Email        string
	PasswordHash string
	LastName     string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
