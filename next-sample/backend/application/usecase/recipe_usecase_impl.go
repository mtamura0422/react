package usecase

import (
	"context"
	"fmt"

	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	sPort "github.com/react/next-sample/backend/domain/service/port"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase/dto"
	"github.com/react/next-sample/backend/usecase/port"
)

var _ port.RecipeUsecase = (*RecipeUsecaseImpl)(nil)

type RecipeUsecaseImpl struct {
	recipeRepository repositories.RecipeRepository
	recipeService    sPort.RecipeService
}

func NewRecipeUsecase(
	recipeRepository repositories.RecipeRepository,
	recipeService sPort.RecipeService,
) port.RecipeUsecase {
	return &RecipeUsecaseImpl{
		recipeRepository: recipeRepository,
		recipeService:    recipeService,
	}
}

func (u *RecipeUsecaseImpl) AddRecipe(
	ctx context.Context,
	recipe *dto.Recipe,
) (*dto.Recipe, *pkgErr.ApplicationError) {

	eRecipe, err := u.recipeService.CreateRecipeTx(ctx, recipe.ToEntity())

	//entity, err := u.recipeRepository.Create(ctx, recipe.ToEntity())
	if err != nil {
		return nil, err
	}

	res, ok := eRecipe.(*entity.Recipe)
	if ok {
		fmt.Println(recipe)
	}

	return dto.ToRecipeMapper(res), nil
}

func (u *RecipeUsecaseImpl) FindRecipe(
	ctx context.Context,
	id int64,
) (*dto.Recipe, *pkgErr.ApplicationError) {
	entity, err := u.recipeRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToRecipeMapper(entity), nil
}

func (u *RecipeUsecaseImpl) GetRecipeList(
	ctx context.Context,
	page int64,
) ([]*dto.Recipe, *pkgErr.ApplicationError) {

	entities, err := u.recipeRepository.List(ctx, page)
	if err != nil {
		return nil, err
	}

	dtoRecipes := make([]*dto.Recipe, len(entities))
	for i, recipeEntity := range entities {
		dtoRecipes[i] = dto.ToRecipeMapper(recipeEntity)
	}

	return dtoRecipes, nil
}
