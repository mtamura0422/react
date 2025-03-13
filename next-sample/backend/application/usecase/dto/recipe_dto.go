package dto

import (
	"time"

	"github.com/react/next-sample/backend/domain/entity"
)

// DTO
type Recipe struct {
	Id              int64
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	RecipeMaterials []RecipeMaterial
}

// DTO モデルを Entity に変換
func (r *Recipe) ToEntity() *entity.Recipe {

	recipeMaterials := make([]entity.RecipeMaterial, len(r.RecipeMaterials))

	for i, rm := range r.RecipeMaterials {
		recipeMaterials[i] = entity.RecipeMaterial{
			Name: rm.Name,
		}
	}
	return &entity.Recipe{
		Id:              r.Id,
		Title:           r.Title,
		Content:         r.Content,
		RecipeMaterials: recipeMaterials,
	}
}

// Entity を DTO(DataTransferObject)に変換
func ToRecipeMapper(e *entity.Recipe) *Recipe {

	recipeMaterial := make([]RecipeMaterial, len(e.RecipeMaterials))

	for i, rm := range e.RecipeMaterials {
		recipeMaterial[i] = *ToRecipeMaterialMapper(&rm)
	}

	return &Recipe{
		Id:              e.Id,
		Title:           e.Title,
		Content:         e.Content,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		RecipeMaterials: recipeMaterial,
	}
}
