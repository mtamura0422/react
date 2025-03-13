package port

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase/dto"
)

type RecipeMaterialUsecase interface {
	FindRecipeMaterial(ctx context.Context, id int64) (*dto.RecipeMaterial, *pkgErr.ApplicationError)
}
