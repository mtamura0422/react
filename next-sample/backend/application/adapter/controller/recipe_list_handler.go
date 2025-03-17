package controller

import (
	"log"
	"net/http"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func (s *Server) GetRecipeList(w http.ResponseWriter, r *http.Request, params openapi.GetRecipeListParams) {
	log.Printf("GetRecipeList handler tuuka1")
	log.Printf("GetRecipeList handler tuuka2 %v", int64(*params.Page))
	recipes, total_count, err := s.recipeUsecase.GetRecipeList(r.Context(), int64(*params.Page))
	if err != nil {
		s.handleError(w, r, err)
		return
	}
	log.Printf("GetRecipeList handler tuuka3")
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
