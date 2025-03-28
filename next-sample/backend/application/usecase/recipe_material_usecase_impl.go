// recipe_material_usecase_impl.go
package usecase

import (
	"context"

	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/usecase/dto"
	"github.com/react/next-sample/backend/usecase/port"
)

var _ port.RecipeMaterialUsecase = (*RecipeMaterialUsecaseImpl)(nil)

// RecipeMaterialUsecaseImpl interface
type RecipeMaterialUsecaseImpl struct {
	recipeMaterialRepository repositories.RecipeMaterialRepository
}

/*
FindRecipeMaterial 材料検索
*/
func (u *RecipeMaterialUsecaseImpl) FindRecipeMaterial(
	ctx context.Context,
	id int64,
) (*dto.RecipeMaterial, error) {
	entity, err := u.recipeMaterialRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToRecipeMaterialMapper(entity), nil
}

// NewRecipeMaterialUsecase create new usecase
func NewRecipeMaterialUsecase(
	recipeMaterialRepository repositories.RecipeMaterialRepository,
) port.RecipeMaterialUsecase {
	return &RecipeMaterialUsecaseImpl{
		recipeMaterialRepository: recipeMaterialRepository,
	}
}
