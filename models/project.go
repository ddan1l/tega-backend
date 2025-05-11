package models

import (
	"database/sql"
	"time"

	project_dto "github.com/ddan1l/tega-backend/dto/project"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Slug        string
	Description string

	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (p *Project) ToDto() *project_dto.ProjectDto {
	if p == nil {
		return nil
	}
	return &project_dto.ProjectDto{
		ID:          p.ID,
		Slug:        p.Slug,
		Name:        p.Name,
		Description: p.Description,
	}
}
