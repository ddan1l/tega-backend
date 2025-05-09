package factory

import (
	"github.com/ddan1l/tega-backend/database"
	"github.com/ddan1l/tega-backend/transaction"
)

type TxManager interface {
	CreateTxManager() database.Database
}

func (f *DefaultFactory) CreateTxManager() transaction.TxManager {
	return transaction.NewTxManager(f.db)

}
