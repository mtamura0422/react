package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe *entity.Recipe) (*entity.Recipe, error)
	Get(ctx context.Context, id int64) (*entity.Recipe, error)
	List(ctx context.Context, page int64) ([]*entity.Recipe, error)
}
