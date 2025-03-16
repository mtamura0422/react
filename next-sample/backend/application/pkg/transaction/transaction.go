package transaction

import (
	"context"
)

type Transaction interface {
	RunInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
