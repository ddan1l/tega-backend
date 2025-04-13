package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model

	Id     int `gorm:"primaryKey;autoIncrement"`
	Token  string
	UserId int
	User   User

	ExpiresAt time.Time `gorm:"default:current_timestamp"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
