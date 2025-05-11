package models

import (
	"database/sql"
	"time"

	user_dto "github.com/ddan1l/tega-backend/dto/user"
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

func (u *User) ToDto() *user_dto.UserDto {
	if u == nil {
		return nil
	}
	return &user_dto.UserDto{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
	}
}
