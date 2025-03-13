package service

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/domain/service/port"
	"github.com/react/next-sample/backend/infrastructure/db"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

var _ port.RecipeService = (*RecipeServiceImpl)(nil)

type RecipeServiceImpl struct {
	txRepo *db.TxRepository
	rRepo  repositories.RecipeRepository
	rmRepo repositories.RecipeMaterialRepository
}

func NewRecipeService(
	txRepo *db.TxRepository,
	rRepo repositories.RecipeRepository,
	rmRepo repositories.RecipeMaterialRepository,
) port.RecipeService {
	return &RecipeServiceImpl{txRepo: txRepo, rRepo: rRepo, rmRepo: rmRepo}
}

func (s *RecipeServiceImpl) createRecipe(input *entity.Recipe) func(ctx context.Context) (interface{}, error) {

	return func(ctx context.Context) (interface{}, error) {

		eRecipe, err := s.rRepo.Create(ctx, input)
		if err != nil {
			return nil, err
		}

		for _, recipeMaterial := range input.RecipeMaterials {
			recipeMaterial.RecipeId = eRecipe.Id
			_, err := s.rmRepo.Create(ctx, &recipeMaterial)
			if err != nil {
				return nil, err
			}

		}

		return eRecipe, nil
	}
}

func (s *RecipeServiceImpl) CreateRecipeTx(ctx context.Context, input *entity.Recipe) (interface{}, *pkgErr.ApplicationError) {

	v, err := s.txRepo.RunInTx(ctx, s.createRecipe(input))
	if err != nil {
		return v, err
	}

	return v, nil
}
