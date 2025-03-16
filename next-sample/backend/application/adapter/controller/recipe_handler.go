package controller

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetVersion")
	log.Printf("listen: GetVersion")
}

func (s *Server) RegisterRecipe(w http.ResponseWriter, r *http.Request) {
	log.Printf("Aaaaaa")
	/*


		log.Printf("aaaaa%v", materialsStr)


			var materials []dto.RecipeMaterial
			err = json.Unmarshal([]byte(materialsStr), &materials)
			if err != nil {
				http.Error(w, "材料情報のパースに失敗しました", http.StatusBadRequest)
				return
			}
	*/

	/*
			// "file" フィールドから取得
			file, handler, err := r.FormFile("file")
			if err != nil {
				http.Error(w, "ファイルがアップロードされていません", http.StatusBadRequest)
				return
			}

		defer file.Close()

		log.Printf("listen: %v", handler.Filename)
	*/

	/*
		var recipe openapi.Recipe
		if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
			s.handleError(w, r, err)
			return
		}
	*/
	/*
		var recipe = openapi.Recipe{
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			Materials: r.MultipartForm.Value["materials[]"],
		}
	*/
	dtoRecipe, e := ReqToDTO(r)
	if e != nil {
		s.handleError(w, r, e)
		return
	}

	uRecipe, err := s.recipeUsecase.AddRecipe(r.Context(), dtoRecipe)

	//uRecipe, err := s.recipeUsecase.AddRecipe(r.Context(), ToDTO(&recipe))
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
