package models

import (
	"time"

	"gorm.io/gorm"
)

type Access struct {
	gorm.Model

	Id   int
	Name string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time
}
