package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	sPort "github.com/react/next-sample/backend/domain/service/port"
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
) (*dto.Recipe, error) {

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
) (*dto.Recipe, error) {
	entity, err := u.recipeRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToRecipeMapper(entity), nil
}

func (u *RecipeUsecaseImpl) GetRecipeList(
	ctx context.Context,
	page int64,
) ([]*dto.Recipe, error) {
	log.Printf("GetRecipeList tuuka0")
	entities, err := u.recipeRepository.List(ctx, page)
	if err != nil {
		return nil, err
	}
	log.Printf("GetRecipeList tuuka1")
	dtoRecipes := make([]*dto.Recipe, len(entities))
	for i, recipeEntity := range entities {
		log.Printf("GetRecipeList tuuka2")
		dtoRecipes[i] = dto.ToRecipeMapper(recipeEntity)
	}
	log.Printf("GetRecipeList tuuka3")
	return dtoRecipes, nil
}

func (u *RecipeUsecaseImpl) SearchRecipeList(
	ctx context.Context,
	word string,
	page int64,
) ([]*dto.Recipe, error) {
	log.Printf("GetRecipeList tuuka0")
	entities, err := u.recipeRepository.Search(ctx, word, page)
	if err != nil {
		return nil, err
	}
	log.Printf("GetRecipeList tuuka1")
	dtoRecipes := make([]*dto.Recipe, len(entities))
	for i, recipeEntity := range entities {
		log.Printf("GetRecipeList tuuka2")
		dtoRecipes[i] = dto.ToRecipeMapper(recipeEntity)
	}
	log.Printf("GetRecipeList tuuka3")
	return dtoRecipes, nil
}
