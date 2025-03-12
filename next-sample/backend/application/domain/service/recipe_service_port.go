package service

import (
	"context"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

type RecipeService interface {
	CreateRecipeTx(ctx context.Context, input *openapi.Recipe) (interface{}, error)
}
