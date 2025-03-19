package port

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

//go:generate mockgen -source=./recipe_service_port.go -destination=./mock/recipe_service_port_mock.go -package=port

type RecipeService interface {
	CreateRecipeTx(ctx context.Context, input *entity.Recipe) (interface{}, error)
}
