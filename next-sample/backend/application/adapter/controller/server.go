// sercer.go
package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	sPort "github.com/react/next-sample/backend/domain/service/port"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase/port"
	"github.com/uptrace/bun"
)

// Server Interface
type Server struct {
	db                    *bun.DB                    // bunインスタンス
	recipeUsecase         port.RecipeUsecase         // recipe usecaseのインターフェース
	recipeMaterialUsecase port.RecipeMaterialUsecase // recipeMaterial usecaseのインターフェース
	recipeService         sPort.RecipeService        // recipe serviceインターフェース
}

// NewServer create new controller
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

		// DI使わないバージョン
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

/*
HandleOK　ハンドルOK
*/
func (s *Server) HandleOK(w http.ResponseWriter, obj interface{}) {
	s.setResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)

	log.Printf("%v", obj)
}

/*
handleError　ハンドルエラー
*/
func (s *Server) handleError(w http.ResponseWriter, r *http.Request, err error) {
	s.setResponseHeaders(w)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))

	log.Printf("%v", r)
	log.Printf("%v", err)

	// 独自エラーならステータスコード表示
	if appErr, ok := err.(*pkgErr.ApplicationError); ok {
		json.NewEncoder(w).Encode(map[string]string{
			"error": appErr.Error(),
			"code":  strconv.Itoa(int(appErr.Code())),
		})

	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
			"code":  "pkg.CodeInternalServerError",
		})
	}

}

/*
setResponseHeaders　レスポンスヘッダをセット
*/
func (s *Server) setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
