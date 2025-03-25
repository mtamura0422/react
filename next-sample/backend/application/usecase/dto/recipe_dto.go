package dto

import (
	"mime/multipart"
	"time"

	"github.com/react/next-sample/backend/domain/entity"
)

// DTO
type Recipe struct {
	Id              int64
	Title           string
	Content         string
	Image           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	RecipeMaterials []RecipeMaterial
	RecipeImage     RecipeImage
}

type RecipeImage struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

const IMAGE_DMAIN = "http://localhost:9000"
const IMAGE_PATH = "/images/"
const IMAGE_DEFAULT_FILW = "sample.png"

// DTO モデルを Entity に変換
func (r *Recipe) ToEntity() *entity.Recipe {

	recipeMaterials := make([]entity.RecipeMaterial, len(r.RecipeMaterials))

	for i, rm := range r.RecipeMaterials {
		recipeMaterials[i] = entity.RecipeMaterial{
			Name: rm.Name,
		}
	}

	RecipeImage := entity.RecipeImage{
		File:       r.RecipeImage.File,
		FileHeader: r.RecipeImage.FileHeader,
	}

	return &entity.Recipe{
		Id:              r.Id,
		Title:           r.Title,
		Content:         r.Content,
		RecipeMaterials: recipeMaterials,
		RecipeImage:     RecipeImage,
	}
}

// Entity を DTO(DataTransferObject)に変換
func ToRecipeMapper(e *entity.Recipe) *Recipe {

	recipeMaterial := make([]RecipeMaterial, len(e.RecipeMaterials))

	for i, rm := range e.RecipeMaterials {
		recipeMaterial[i] = *ToRecipeMaterialMapper(&rm)
	}

	filename := IMAGE_DMAIN + IMAGE_PATH

	if e.Image != "" {
		filename += e.Image
	} else {
		filename += IMAGE_DEFAULT_FILW
	}

	return &Recipe{
		Id:              e.Id,
		Title:           e.Title,
		Content:         e.Content,
		Image:           filename,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		RecipeMaterials: recipeMaterial,
	}
}
