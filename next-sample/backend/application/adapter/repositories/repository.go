package repositories

import (
	"database/sql"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

func RepositoryError(err error) *pkgErr.ApplicationError {
	switch err {
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
