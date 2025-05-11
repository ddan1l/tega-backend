package models

import (
	"database/sql"
	"time"

	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"gorm.io/gorm"
)

type ProjectUser struct {
	gorm.Model

	ID int `gorm:"primaryKey;autoIncrement"`

	UserID int
	User   User `gorm:"foreignKey:UserID"`

	ProjectID int
	Project   Project `gorm:"foreignKey:ProjectID"`

	RoleID int
	Role   Role `gorm:"foreignKey:RoleID"`

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (pu *ProjectUser) ToDto() *project_dto.ProjectUserDto {
	return &project_dto.ProjectUserDto{
		ID:        pu.ID,
		UserID:    pu.UserID,
		RoleID:    pu.RoleID,
		ProjectID: pu.ProjectID,
		User:      pu.User.ToDto(),
		Project:   pu.Project.ToDto(),
	}
}
