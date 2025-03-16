package port

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

type RecipeService interface {
	CreateRecipeTx(ctx context.Context, input *entity.Recipe) (interface{}, error)
}
