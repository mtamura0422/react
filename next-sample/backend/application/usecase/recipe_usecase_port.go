package usecase

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeUsecase interface {
	AddRecipe(ctx context.Context, recipe *Recipe) (*Recipe, *pkgErr.ApplicationError)
	FindRecipe(ctx context.Context, id int64) (*Recipe, *pkgErr.ApplicationError)
	GetRecipeList(ctx context.Context, page int64) ([]*Recipe, *pkgErr.ApplicationError)
}
