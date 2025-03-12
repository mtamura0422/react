package repositories

import (
	"database/sql"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type TxKey string

const TX_KEY TxKey = "TRANSACTION_KEY"

func RepositoryError(err error) *pkgErr.ApplicationError {
	switch err {
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
