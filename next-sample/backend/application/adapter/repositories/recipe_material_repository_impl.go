package repositories

import (
	"context"

	"github.com/react/next-sample/backend/adapter/repositories/model"
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/uptrace/bun"
)

var _ repositories.RecipeMaterialRepository = (*RecipeMaterialRepositoryImpl)(nil)

type RecipeMaterialRepositoryImpl struct {
	db *bun.DB
}

func (r *RecipeMaterialRepositoryImpl) Create(ctx context.Context, eRecipeMaterial *entity.RecipeMaterial) (*entity.RecipeMaterial, error) {

	var inserter *bun.InsertQuery
	var recipeMaterial *model.RecipeMaterial

	recipeMaterial = model.ToRecipeMaterialMapper(eRecipeMaterial)

	// context からトランザクションオブジェクトを取得する
	if tx, ok := db.GetTx(ctx); ok {
		// トランザクションオブジェクトがある場合は、トランザクションで処理を行う
		inserter = tx.NewInsert()
	} else {
		// トランザクションオブジェクトがない場合は通常の処理を行う
		inserter = r.db.NewInsert()
	}

	_, err := inserter.Model(recipeMaterial).Exec(ctx)
	if err != nil {
		return nil, RepositoryError(err)
	}
	return recipeMaterial.ToEntity(), nil

}

func (u *RecipeMaterialRepositoryImpl) Get(ctx context.Context, id int64) (*entity.RecipeMaterial, error) {
	var recipeMaterial model.RecipeMaterial
	if err := u.db.NewSelect().Model(&recipeMaterial).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return recipeMaterial.ToEntity(), nil
}

func NewRecipeMaterialRepository(db *bun.DB) repositories.RecipeMaterialRepository {
	return &RecipeMaterialRepositoryImpl{db: db}
}
