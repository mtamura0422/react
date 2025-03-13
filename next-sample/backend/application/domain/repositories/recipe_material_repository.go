package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeMaterialRepository interface {
	Get(ctx context.Context, id int64) (*entity.RecipeMaterial, *pkgErr.ApplicationError)
	Create(ctx context.Context, recipe *entity.RecipeMaterial) (*entity.RecipeMaterial, *pkgErr.ApplicationError)
}
