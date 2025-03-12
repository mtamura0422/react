package usecase

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

type RecipeMaterialUsecase interface {
	FindRecipeMaterial(ctx context.Context, id int64) (*RecipeMaterial, *pkgErr.ApplicationError)
}
