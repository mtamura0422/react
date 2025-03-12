package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetVersion")
	log.Printf("listen: GetVersion")
}

func (s *Server) RegisterRecipe(w http.ResponseWriter, r *http.Request) {

	log.Printf("listen: RegisterRecipe")
	var recipe openapi.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		s.handleError(w, r, err)
		return
	}
	log.Printf("listen: RegisterRecipe2")
	log.Printf("%v", &recipe)

	uRecipe, err := s.recipeService.CreateRecipeTx(r.Context(), &recipe)

	res, ok := uRecipe.(*openapi.Recipe)
	if ok {
		fmt.Println(recipe)
	}

	log.Printf("あああああああああ")
	//uRecipe, err := s.recipeUsecase.AddRecipe(r.Context(), ToDTO(&recipe))
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	s.HandleOK(w, res)

}

func (s *Server) GetRecipe(w http.ResponseWriter, r *http.Request, recipeId int64) {

	recipe, err := s.recipeUsecase.FindRecipe(r.Context(), recipeId)
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	s.HandleOK(w, ToResponse(recipe))

}
