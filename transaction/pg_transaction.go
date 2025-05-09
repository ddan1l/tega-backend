package transaction

import (
	"errors"
	"log"

	"github.com/ddan1l/tega-backend/database"
	errs "github.com/ddan1l/tega-backend/errors"
)

type txManager struct {
	db database.Database
	tx database.Database
}

func NewTxManager(db database.Database) TxManager {
	return &txManager{db: db}
}

func (m *txManager) Begin() error {
	if m.tx != nil {
		return errors.New("transaction already started")
	}
	m.tx = database.NewTransaction(m.db)
	return nil
}

func (m *txManager) Commit() error {
	return m.tx.GetDb().Commit().Error
}

func (m *txManager) GetTx() database.Database {
	return m.tx
}

func (m *txManager) Rollback() error {
	if m.tx == nil {
		return nil
	}
	err := m.tx.GetDb().Rollback().Error
	m.tx = nil
	return err
}

func (m *txManager) CallWithTx(fn func(tx database.Database) *errs.AppError) *errs.AppError {
	log.Println("\n\n====== BEGIN TRANSACTION ======")

	tx := database.NewTransaction(m.db)

	defer func() {
		log.Println("\n\n====== ROLLBACK TRANSACTION ======")
		if r := recover(); r != nil {
			tx.GetDb().Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		log.Println("\n\n====== ROLLBACK TRANSACTION ======")
		tx.GetDb().Rollback()
		return err
	}

	if err := tx.GetDb().Commit().Error; err != nil {
		log.Println("\n\n====== ROLLBACK TRANSACTION ======")
		tx.GetDb().Rollback()
		return errs.BadRequest.WithError(err)
	}

	log.Println("\n\n====== COMMIT TRANSACTION ======")
	return nil
}
