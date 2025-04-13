package factory

import (
	"github.com/ddan1l/tega-backend/database"
)

type DefaultFactory struct {
	db database.Database
}

func NewDefaultFactory(db database.Database) *DefaultFactory {
	return &DefaultFactory{db: db}
}
