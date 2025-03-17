package repositories

import (
	"context"
	"log"

	"github.com/react/next-sample/backend/adapter/repositories/model"
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/uptrace/bun"
)

var _ repositories.RecipeRepository = (*RecipeRepositoryImpl)(nil)

type RecipeRepositoryImpl struct {
	db *bun.DB
}

const LIMIT = 10

func (r *RecipeRepositoryImpl) Create(ctx context.Context, eRecipe *entity.Recipe) (*entity.Recipe, error) {

	var inserter *bun.InsertQuery
	var recipe *model.Recipe

	recipe = model.ToRecipeMapper(eRecipe)

	// context からトランザクションオブジェクトを取得する
	if tx, ok := db.GetTx(ctx); ok {
		// トランザクションオブジェクトがある場合は、トランザクションで処理を行う
		inserter = tx.NewInsert()
	} else {
		// トランザクションオブジェクトがない場合は通常の処理を行う
		inserter = r.db.NewInsert()
	}

	ret, err := inserter.Model(recipe).Exec(ctx)
	if err != nil {
		return nil, RepositoryError(err)
	}

	lastInsertID, err := ret.LastInsertId()

	recipe.Id = lastInsertID
	log.Printf("recipe Create tuuka1")
	return recipe.ToEntity(), nil

}

func (u *RecipeRepositoryImpl) Get(ctx context.Context, id int64) (*entity.Recipe, error) {

	var recipe model.Recipe
	if err := u.db.NewSelect().Model(&recipe).Relation("RecipeMaterials").Where("id = ?", id).Scan(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return recipe.ToEntity(), nil
}

func (u *RecipeRepositoryImpl) List(ctx context.Context, page int64) ([]*entity.Recipe, int, error) {

	var dbRecipes []model.Recipe

	offset := int((page - 1) * LIMIT)

	count, err := u.db.NewSelect().Model(&dbRecipes).
		Relation("RecipeMaterials").
		Offset(offset).
		Order("id desc").
		Limit(LIMIT).
		ScanAndCount(ctx)

	if err != nil {
		return nil, 0, RepositoryError(err)
	}

	recipeEntityes := make([]*entity.Recipe, len(dbRecipes))
	for i, recipeRecord := range dbRecipes {
		recipeEntityes[i] = recipeRecord.ToEntity()
	}

	return recipeEntityes, count, nil
}

func (u *RecipeRepositoryImpl) Search(ctx context.Context, word string, page int64) ([]*entity.Recipe, int, error) {

	var dbRecipes []model.Recipe

	offset := int((page - 1) * LIMIT)

	count, err := u.db.NewSelect().Model(&dbRecipes).
		Relation("RecipeMaterials").
		Where("(title LIKE ?) or (content LIKE ?)", "%"+word+"%", "%"+word+"%").
		Offset(offset).
		Order("id desc").
		Limit(LIMIT).
		ScanAndCount(ctx)

	if err != nil {
		return nil, 0, RepositoryError(err)
	}

	recipeEntityes := make([]*entity.Recipe, len(dbRecipes))
	for i, recipeRecord := range dbRecipes {
		recipeEntityes[i] = recipeRecord.ToEntity()
	}

	return recipeEntityes, count, nil
}

func NewRecipeRepository(db *bun.DB) repositories.RecipeRepository {
	return &RecipeRepositoryImpl{db: db}
}
