package usecase

import (
	"context"
	"log"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

var _ RecipeMaterialUsecase = (*RecipeMaterialUsecaseImpl)(nil)

type RecipeMaterialUsecaseImpl struct {
	recipeMaterialRepository RecipeMaterialRepository
}

func (u *RecipeMaterialUsecaseImpl) FindRecipeMaterial(
	ctx context.Context,
	id int64,
) (*RecipeMaterial, *pkgErr.ApplicationError) {
	log.Printf("tuuka material2")
	entity, err := u.recipeMaterialRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	log.Printf("tuuka material3")
	return ToRecipeMaterialMapper(entity), nil
}

func NewRecipeMaterialUsecase(
	recipeMaterialRepository RecipeMaterialRepository,
) RecipeMaterialUsecase {
	return &RecipeMaterialUsecaseImpl{
		recipeMaterialRepository: recipeMaterialRepository,
	}
}
