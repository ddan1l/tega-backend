package transaction

import (
	"github.com/ddan1l/tega-backend/database"
	errs "github.com/ddan1l/tega-backend/errors"
)

type TxManager interface {
	Begin() error
	Commit() error
	Rollback() error
	GetTx() database.Database
	CallWithTx(fn func(tx database.Database) *errs.AppError) *errs.AppError
}
