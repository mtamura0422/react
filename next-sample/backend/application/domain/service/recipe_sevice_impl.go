package service

import (
	"context"
	"log"

	"github.com/react/next-sample/backend/adapter/repositories"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/react/next-sample/backend/infrastructure/openapi"
)

var _ RecipeService = (*RecipeServiceImpl)(nil)

type RecipeServiceImpl struct {
	txRepo *db.TxRepository
	rRepo  repositories.RecipeRepositoryImpl
	rmRepo repositories.RecipeaMaterialRepositoryImpl
}

func NewRecipeService(
	txRepo *db.TxRepository,
	rRepo repositories.RecipeRepositoryImpl,
	rmRepo repositories.RecipeaMaterialRepositoryImpl,
) *RecipeServiceImpl {
	return &RecipeServiceImpl{txRepo: txRepo, rRepo: rRepo, rmRepo: rmRepo}
}

func (s *RecipeServiceImpl) createRecipe(input *openapi.Recipe) func(ctx context.Context) (interface{}, error) {

	return func(ctx context.Context) (interface{}, error) {

		recipe := ToEntity(input)
		eRecipe, err := s.rRepo.Add(ctx, recipe)
		if err != nil {
			return nil, err
		}

		recipeMaterials := ToEntityMaterial(input)

		for _, recipeMaterial := range recipeMaterials {

			recipeMaterial.RecipeId = eRecipe.Id
			_, err := s.rmRepo.Add(ctx, recipeMaterial)
			if err != nil {
				return nil, err
			}

		}

		log.Printf("listen: RegisterRecipe6")
		return ToResponse(eRecipe), nil
	}
}

func (s *RecipeServiceImpl) CreateRecipeTx(ctx context.Context, input *openapi.Recipe) (interface{}, error) {

	v, err := s.txRepo.RunInTx(ctx, s.createRecipe(input))
	if err != nil {
		return v, err
	}
	return v, nil
}
