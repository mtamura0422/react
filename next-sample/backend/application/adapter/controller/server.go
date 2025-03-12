package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/react/next-sample/backend/adapter/repositories"
	"github.com/react/next-sample/backend/domain/service"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/react/next-sample/backend/usecase"
	"github.com/uptrace/bun"
)

// ServerInterfaceを実装
type Server struct {
	db                    *bun.DB
	recipeUsecase         usecase.RecipeUsecase
	recipeMaterialUsecase usecase.RecipeMaterialUsecase
	recipeService         service.RecipeService
}

func NewServer() *Server {

	d, _ := db.NewDB()
	tx := db.NewTxRepository(d)

	rp := repositories.NewRecipeRepository(d)
	rpm := repositories.NewRecipeMaterialRepository(d)
	rr := repositories.NewRecipeRepositoryImpl(d)
	rmr := repositories.NewRecipeaMaterialRepositoryImpl(d)

	return &Server{
		db:                    d,
		recipeUsecase:         usecase.NewRecipeUsecase(rp),
		recipeMaterialUsecase: usecase.NewRecipeMaterialUsecase(rpm),
		recipeService:         service.NewRecipeService(tx, *rr, *rmr),
	}
}

func (s *Server) SetDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.db.RunInTx(r.Context(), nil, func(ctx context.Context, tx bun.Tx) error {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(context.WithValue(r.Context(), repositories.TX_KEY, &tx)))

			if ww.Status() != http.StatusOK {
				return errors.New("Rollbacked") // rollback
			}

			// commit
			return nil
		})
	})
}

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
