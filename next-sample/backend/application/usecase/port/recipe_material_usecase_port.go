package port

import (
	"context"

	"github.com/react/next-sample/backend/usecase/dto"
)

type RecipeMaterialUsecase interface {
	FindRecipeMaterial(ctx context.Context, id int64) (*dto.RecipeMaterial, error)
}
