package controller

import (
	"net/http"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func (s *Server) GetRecipeList(w http.ResponseWriter, r *http.Request, params openapi.GetRecipeListParams) {
	recipes, err := s.recipeUsecase.GetRecipeList(r.Context(), int64(*params.Page))
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	oapiRecipes := make([]*openapi.Recipe, len(recipes))
	for i, recipe := range recipes {
		oapiRecipes[i] = ToResponse(recipe)
	}

	s.HandleOK(w, oapiRecipes)
}
