package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/react/next-sample/backend/infrastructure/openapi"
	pkgErr "github.com/react/next-sample/backend/pkg/error"
	"github.com/react/next-sample/backend/usecase/dto"
)

func ReqToDTO(r *http.Request) (*dto.Recipe, error) {

	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		return nil, err
	}

	title := r.FormValue("title")
	materialsStr := r.FormValue("materials")
	content := r.FormValue("content")

	// 材料

	var materials []string
	err = json.Unmarshal([]byte(materialsStr), &materials)

	if err != nil {
		return nil, pkgErr.NewApplicationError("材料のデコードに失敗しました", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}

	recipeMaterials := make([]dto.RecipeMaterial, len(materials))
	for i, rm := range materials {

		if utf8.RuneCountInString(rm) > 15 {
			return nil, pkgErr.NewApplicationError("材料 15文字以内にしてください", pkgErr.LevelError, pkgErr.CodeValidatorError)
		}

		recipeMaterials[i] = dto.RecipeMaterial{
			Name: rm,
		}
	}

	log.Printf("%v", map[string]string{"error": "test message"})

	if len(materials) <= 0 {
		return nil, pkgErr.NewApplicationError("材料 は必須です", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}

	if title == "" {
		return nil, pkgErr.NewApplicationError("レシピタイトル は必須です", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}
	if content == "" {
		return nil, pkgErr.NewApplicationError("作り方 は必須です", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}
	if utf8.RuneCountInString(title) > 20 {
		return nil, pkgErr.NewApplicationError("レシピタイトル は20文字以内にしてください", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}
	if utf8.RuneCountInString(content) > 500 {
		return nil, pkgErr.NewApplicationError("作り方 は500文字以内にしてください", pkgErr.LevelError, pkgErr.CodeValidatorError)
	}

	file, fileHeader, err := r.FormFile("file")

	var RecipeImage dto.RecipeImage

	if err == nil {
		RecipeImage = dto.RecipeImage{
			File:       file,
			FileHeader: fileHeader,
		}
	}

	return &dto.Recipe{
		Title:           title,
		Content:         content,
		RecipeMaterials: recipeMaterials,
		RecipeImage:     RecipeImage,
	}, nil
}

func ToDTO(req *openapi.Recipe) *dto.Recipe {

	recipeMaterials := make([]dto.RecipeMaterial, len(req.Materials))

	for i, rm := range req.Materials {
		recipeMaterials[i] = dto.RecipeMaterial{
			Name: rm,
		}
	}

	return &dto.Recipe{
		Id:              req.Id,
		Title:           req.Title,
		Content:         req.Content,
		CreatedAt:       req.CreatedAt,
		UpdatedAt:       req.UpdatedAt,
		RecipeMaterials: recipeMaterials,
	}
}

func ToResponse(u *dto.Recipe) *openapi.Recipe {

	materials := make([]string, len(u.RecipeMaterials))

	for i, rm := range u.RecipeMaterials {
		materials[i] = rm.Name
	}

	return &openapi.Recipe{
		Id:        u.Id,
		Title:     u.Title,
		Content:   u.Content,
		Image:     u.Image,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Materials: materials,
	}
}
