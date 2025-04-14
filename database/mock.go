package database

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockDatabase struct {
	Db *gorm.DB
}

var (
	dbMockInstance *mockDatabase
)

func NewMockDatabase() (Database, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	dbMockInstance = &mockDatabase{Db: gormDB}

	return dbMockInstance, mock
}

func (p *mockDatabase) GetDb() *gorm.DB {
	return dbMockInstance.Db
}
