// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	nethttp_middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/react/next-sample/backend/di"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func InitRouter() {

	openapi3filter.RegisterBodyDecoder("multipart/form-data", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/jpeg", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/png", openapi3filter.FileBodyDecoder)

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
	//server := my_controller.NewServer()
	server := di.InitializeServer()
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

	router.Use(middlewareFormdataValidator)

	router.Use(middleware.Heartbeat("/healthz"))
	router.Use(middleware.AllowContentType("application/json", "multipart/form-data"))

	router.Use(oapiRequestValidatorWithExclusion(swagger))
	//router.Use(nethttp_middleware.OapiRequestValidator(swagger)) // validationを設定

	openapi.HandlerFromMux(server, router) // chiのrouterと実装したserverを紐付け
	if err := http.ListenAndServe(":9000", router); err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %s\n", err)
		os.Exit(1)
	}

	db.CloseDB()

}

// 画像配信用
var middlewareFormdataValidator = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// `multipart/form-data` の場合は `OapiRequestValidator` の前にリクエストを解析する
		if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			err := r.ParseMultipartForm(10 << 20) // 10MB
			if err != nil {
				http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

var middlewareStaticImages = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			len := r.ContentLength
			body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
			r.Body.Read(body)         // byte 配列にリクエストボディを読み込む

			log.Printf("body ====")
			log.Println(w, string(body))
		*/
		ret := regexp.MustCompile("^/images/.*")

		if ret.MatchString(r.URL.Path) {
			fileServer := http.FileServer(http.Dir("images/"))
			http.StripPrefix("/images/", fileServer).ServeHTTP(w, r)
			//	w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func oapiRequestValidatorWithExclusion(swagger *openapi3.T) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// `multipart/form-data` の場合は `OapiRequestValidator` の前にリクエストを解析する
			if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
				err := r.ParseMultipartForm(10 << 20) // 10MB
				if err != nil {
					http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
					return
				}
				next.ServeHTTP(w, r)
			} else {
				validator := nethttp_middleware.OapiRequestValidator(swagger)
				validator(next).ServeHTTP(w, r)
			}
		})
	}

	/*
		validator := nethttp_middleware.OapiRequestValidator(swagger)
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/recipe/register" {
					next.ServeHTTP(w, r)
					return
				}
				validator(next).ServeHTTP(w, r)
			})
		}
	*/
}
