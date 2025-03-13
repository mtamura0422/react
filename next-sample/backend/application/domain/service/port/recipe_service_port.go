package port

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeService interface {
	CreateRecipeTx(ctx context.Context, input *entity.Recipe) (interface{}, *pkgErr.ApplicationError)
}
