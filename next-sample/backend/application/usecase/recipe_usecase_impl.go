package usecase

import (
	"context"

	pkgErr "github.com/react/next-sample/backend/pkg/error"
)

var _ RecipeUsecase = (*RecipeUsecaseImpl)(nil)

type RecipeUsecaseImpl struct {
	recipeRepository RecipeRepository
}

func (u *RecipeUsecaseImpl) AddRecipe(
	ctx context.Context,
	recipe *Recipe,
) (*Recipe, *pkgErr.ApplicationError) {
	entity, err := u.recipeRepository.Add(ctx, recipe.ToEntity())
	if err != nil {
		return nil, err
	}

	return ToRecipeMapper(entity), nil
}

func (u *RecipeUsecaseImpl) FindRecipe(
	ctx context.Context,
	id int64,
) (*Recipe, *pkgErr.ApplicationError) {
	entity, err := u.recipeRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToRecipeMapper(entity), nil
}

func (u *RecipeUsecaseImpl) GetRecipeList(
	ctx context.Context,
	page int64,
) ([]*Recipe, *pkgErr.ApplicationError) {

	entities, err := u.recipeRepository.GetList(ctx, page)
	if err != nil {
		return nil, err
	}

	dtoRecipes := make([]*Recipe, len(entities))
	for i, recipeEntity := range entities {
		dtoRecipes[i] = ToRecipeMapper(recipeEntity)
	}

	return dtoRecipes, nil
}

func NewRecipeUsecase(
	recipeRepository RecipeRepository,
) RecipeUsecase {
	return &RecipeUsecaseImpl{
		recipeRepository: recipeRepository,
	}
}
