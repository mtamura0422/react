package controller

import (
	"encoding/json"
	"log"
	"net/http"

	sPort "github.com/react/next-sample/backend/domain/service/port"
	"github.com/react/next-sample/backend/usecase/port"
	"github.com/uptrace/bun"
)

// ServerInterfaceを実装
type Server struct {
	db                    *bun.DB
	recipeUsecase         port.RecipeUsecase
	recipeMaterialUsecase port.RecipeMaterialUsecase
	recipeService         sPort.RecipeService
}

func NewServer(
	db *bun.DB,
	recipeUsecase port.RecipeUsecase,
	recipeMaterialUsecase port.RecipeMaterialUsecase,
	recipeService sPort.RecipeService,

) *Server {
	return &Server{
		db:                    db,
		recipeUsecase:         recipeUsecase,
		recipeMaterialUsecase: recipeMaterialUsecase,
		recipeService:         recipeService,
	}
}

/*
func NewServer() *Server {

		d, _ := db.NewDB()
		tx := db.NewTxRepository(d)

		rp := repositories.NewRecipeRepository(d)
		rpm := repositories.NewRecipeMaterialRepository(d)
		rs := service.NewRecipeService(tx, rp, rpm)

	return &Server{
		db:                    d,
		recipeUsecase:         usecase.NewRecipeUsecase(rp, rs),
		recipeMaterialUsecase: usecase.NewRecipeMaterialUsecase(rpm),
		recipeService:         service.NewRecipeService(tx, rp, rpm),
	}
}
*/

func (s *Server) HandleOK(w http.ResponseWriter, obj interface{}) {
	s.setResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func (s *Server) handleError(w http.ResponseWriter, r *http.Request, err error) {
	s.setResponseHeaders(w)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

	log.Printf("%v", r)
	log.Printf("%v", err)

}

func (s *Server) setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
