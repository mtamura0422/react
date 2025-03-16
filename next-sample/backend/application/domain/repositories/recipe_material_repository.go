package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

type RecipeMaterialRepository interface {
	Get(ctx context.Context, id int64) (*entity.RecipeMaterial, error)
	Create(ctx context.Context, recipe *entity.RecipeMaterial) (*entity.RecipeMaterial, error)
}
