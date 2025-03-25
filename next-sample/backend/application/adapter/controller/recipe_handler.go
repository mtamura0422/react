package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetVersion")
	log.Printf("listen: GetVersion")
}

func (s *Server) GetRecipeSearch(w http.ResponseWriter, r *http.Request, params openapi.GetRecipeSearchParams) {
	recipes, total_count, err := s.recipeUsecase.SearchRecipeList(r.Context(), params.Q, int64(*params.Page))
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

func (s *Server) RegisterRecipe(w http.ResponseWriter, r *http.Request) {

	dtoRecipe, e := ReqToDTO(r)
	if e != nil {
		s.handleError(w, r, e)
		return
	}

	uRecipe, err := s.recipeUsecase.AddRecipe(r.Context(), dtoRecipe)

	if err != nil {
		s.handleError(w, r, err)
		return
	}

	s.HandleOK(w, ToResponse(uRecipe))

}

func (s *Server) GetRecipe(w http.ResponseWriter, r *http.Request, recipeId int64) {

	recipe, err := s.recipeUsecase.FindRecipe(r.Context(), recipeId)
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	s.HandleOK(w, ToResponse(recipe))

}
