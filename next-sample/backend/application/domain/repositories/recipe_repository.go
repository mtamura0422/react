package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

//go:generate mockgen -source=./recipe_repository.go -destination=./mock/recipe_repository_mock.go -package=repositories

type RecipeRepository interface {
	Create(ctx context.Context, recipe *entity.Recipe) (*entity.Recipe, error)
	Get(ctx context.Context, id int64) (*entity.Recipe, error)
	List(ctx context.Context, page int64) ([]*entity.Recipe, int, error)
	Search(ctx context.Context, word string, page int64) ([]*entity.Recipe, int, error)
}
