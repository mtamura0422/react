package port

import (
	"context"

	"github.com/react/next-sample/backend/usecase/dto"
)

type RecipeUsecase interface {
	AddRecipe(ctx context.Context, recipe *dto.Recipe) (*dto.Recipe, error)
	FindRecipe(ctx context.Context, id int64) (*dto.Recipe, error)
	GetRecipeList(ctx context.Context, page int64) ([]*dto.Recipe, int, error)
	SearchRecipeList(ctx context.Context, word string, page int64) ([]*dto.Recipe, int, error)
}
