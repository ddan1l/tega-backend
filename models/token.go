package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model

	ID    int `gorm:"primaryKey;autoIncrement"`
	Token string

	UserID int
	User   User `gorm:"foreignKey:UserID"`

	ExpiresAt time.Time `gorm:"default:current_timestamp"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
