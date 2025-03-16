package usecase

import (
	"context"
	"log"

	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/usecase/dto"
	"github.com/react/next-sample/backend/usecase/port"
)

var _ port.RecipeMaterialUsecase = (*RecipeMaterialUsecaseImpl)(nil)

type RecipeMaterialUsecaseImpl struct {
	recipeMaterialRepository repositories.RecipeMaterialRepository
}

func (u *RecipeMaterialUsecaseImpl) FindRecipeMaterial(
	ctx context.Context,
	id int64,
) (*dto.RecipeMaterial, error) {
	log.Printf("tuuka material2")
	entity, err := u.recipeMaterialRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	log.Printf("tuuka material3")
	return dto.ToRecipeMaterialMapper(entity), nil
}

func NewRecipeMaterialUsecase(
	recipeMaterialRepository repositories.RecipeMaterialRepository,
) port.RecipeMaterialUsecase {
	return &RecipeMaterialUsecaseImpl{
		recipeMaterialRepository: recipeMaterialRepository,
	}
}
