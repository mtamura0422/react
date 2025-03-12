package usecase

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeMaterialRepository interface {
	Find(ctx context.Context, id int64) (*entity.RecipeMaterial, *pkgErr.ApplicationError)
}
