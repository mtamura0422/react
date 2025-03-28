// recipe_list_handler.go
package controller

import (
	"net/http"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

/*
GetRecipeList　レシピリスト
*/
func (s *Server) GetRecipeList(w http.ResponseWriter, r *http.Request, params openapi.GetRecipeListParams) {

	recipes, total_count, err := s.recipeUsecase.GetRecipeList(r.Context(), int64(*params.Page))
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	oapiRecipes := make([]openapi.Recipe, len(recipes))
	for i, recipe := range recipes {
		oapiRecipes[i] = *ToResponse(recipe)
	}

	recipeList := openapi.RecipeList{
		List:       oapiRecipes,
		Page:       int(*params.Page),
		PerPage:    10,
		TotalCount: total_count,
	}

	s.HandleOK(w, recipeList)
}
