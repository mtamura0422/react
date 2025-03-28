// repository.go
package repositories

import (
	"database/sql"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

/*
RepositoryError repositories用エラーオブジェクト生成
*/
func RepositoryError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
