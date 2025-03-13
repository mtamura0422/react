package transaction

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type Transaction interface {
	RunInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, *pkgErr.ApplicationError)
}
