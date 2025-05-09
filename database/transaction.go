package database

import (
	"gorm.io/gorm"
)

type postgresTransaction struct {
	Db *gorm.DB
}

func NewTransaction(db Database) Database {
	return &postgresTransaction{
		Db: db.GetDb().Begin(),
	}
}

func (p *postgresTransaction) GetDb() *gorm.DB {
	return p.Db
}
