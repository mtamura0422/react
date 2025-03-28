// transaction.go
package db

import (
	"context"
	"log"

	"github.com/uptrace/bun"

	"database/sql"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/pkg/transaction"
)

var _ transaction.Transaction = (*TxRepository)(nil)

type txKey struct{}

// トランザクション保存用のcontextのkey
var TxCtxKey = txKey{}

// TxRepository Interface
type TxRepository struct {
	db *bun.DB
}

// GetDBConn はTxRepositoryが保持しているConnectionを返します．
func (tr *TxRepository) GetDBConn() *bun.DB {
	return tr.db
}

func NewTxRepository(db *bun.DB) *TxRepository {
	return &TxRepository{db: db}
}

/*
RunInTx トランザクション処理実行
*/
func (r *TxRepository) RunInTx(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	log.Printf("listen: RegisterRecipe7")
	log.Printf("listen: tuuka")
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("listen: err")
		return nil, TransactionError(err)
	}
	log.Printf("listen: RegisterRecipe8")
	// bunの `RunInTx`をベースに途中でcontextにトランザクションオブジェクトを入れる処理を追加
	c := context.WithValue(ctx, TxCtxKey, tx)

	var done bool

	defer func() {
		if !done {
			_ = tx.Rollback()
		}
	}()

	v, err := fn(c)
	if err != nil {
		return v, TransactionError(err)
	}

	done = true
	return v, TransactionError(tx.Commit())
}

/*
GetTx context.Contextからトランザクションを取得
*/
func GetTx(ctx context.Context) (*bun.Tx, bool) {
	tx, ok := ctx.Value(TxCtxKey).(*bun.Tx)
	return tx, ok
}

/*
TransactionError エラーオブジェクト生成
*/
func TransactionError(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
