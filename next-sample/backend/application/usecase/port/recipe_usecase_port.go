package port

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase/dto"
)

type RecipeUsecase interface {
	AddRecipe(ctx context.Context, recipe *dto.Recipe) (*dto.Recipe, *pkgErr.ApplicationError)
	FindRecipe(ctx context.Context, id int64) (*dto.Recipe, *pkgErr.ApplicationError)
	GetRecipeList(ctx context.Context, page int64) ([]*dto.Recipe, *pkgErr.ApplicationError)
}
