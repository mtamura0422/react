package dto

import (
	"time"

	"github.com/react/next-sample/backend/domain/entity"
)

// DTO
type RecipeMaterial struct {
	Id        int64
	RecipeId  int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Entity を ORM モデルに変換
func ToRecipeMaterialMapper(e *entity.RecipeMaterial) *RecipeMaterial {
	return &RecipeMaterial{
		Id:        e.Id,
		RecipeId:  e.RecipeId,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
