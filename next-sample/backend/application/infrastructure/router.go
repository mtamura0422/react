// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	nethttp_middmiddleware "github.com/oapi-codegen/nethttp-middleware"
	my_controller "github.com/react/next-sample/backend/adapter/controller"
	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func InitRouter() {
	swagger, err := openapi.GetSwagger() // APIスキーマ定義を取得
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	// chiのログ出力設定
	logger := httplog.NewLogger("server", httplog.Options{
		JSON: true,
	})
	server := my_controller.NewServer()
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(httplog.RequestLogger(logger))
	router.Use(middleware.Recoverer)

	router.Use(middlewareStaticImages)
	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(nethttp_middmiddleware.OapiRequestValidator(swagger)) // validationを設定

	openapi.HandlerFromMux(server, router) // chiのrouterと実装したserverを紐付け
	http.ListenAndServe(":9000", router)

}

// 画像配信用
var middlewareStaticImages = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ret := regexp.MustCompile("^/images/*")

		if ret.MatchString(r.URL.Path) {
			fileServer := http.FileServer(http.Dir("images/"))
			http.StripPrefix("/images/", fileServer).ServeHTTP(w, r)
			//	w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
