package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe *entity.Recipe) (*entity.Recipe, *pkgErr.ApplicationError)
	Get(ctx context.Context, id int64) (*entity.Recipe, *pkgErr.ApplicationError)
	List(ctx context.Context, page int64) ([]*entity.Recipe, *pkgErr.ApplicationError)
}
