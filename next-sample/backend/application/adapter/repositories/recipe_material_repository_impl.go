package repositories

import (
	"context"
	"log"

	"github.com/react/next-sample/backend/adapter/repositories/model"
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/infrastructure/db"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase"
	"github.com/uptrace/bun"
)

var _ usecase.RecipeMaterialRepository = (*RecipeaMaterialRepositoryImpl)(nil)

type RecipeaMaterialRepositoryImpl struct {
	db *bun.DB
}

func (r *RecipeaMaterialRepositoryImpl) Add(ctx context.Context, eRecipeMaterial *entity.RecipeMaterial) (*entity.RecipeMaterial, *pkgErr.ApplicationError) {

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

func (u *RecipeaMaterialRepositoryImpl) Find(ctx context.Context, id int64) (*entity.RecipeMaterial, *pkgErr.ApplicationError) {
	log.Printf("tuuka1")

	//tx := ctx.Value(TX_KEY).(*bun.Tx)
	log.Printf("tuuka2")
	var recipeMaterial model.RecipeMaterial
	if err := u.db.NewSelect().Model(&recipeMaterial).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return recipeMaterial.ToEntity(), nil
}

func NewRecipeMaterialRepository(db *bun.DB) usecase.RecipeMaterialRepository {
	return &RecipeaMaterialRepositoryImpl{db: db}
}

func NewRecipeaMaterialRepositoryImpl(db *bun.DB) *RecipeaMaterialRepositoryImpl {
	return &RecipeaMaterialRepositoryImpl{db: db}
}
