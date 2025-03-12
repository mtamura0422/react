package usecase

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

// ORM モデルを Entity に変換
func (rm *RecipeMaterial) ToEntity() *entity.RecipeMaterial {
	return &entity.RecipeMaterial{
		Id:       rm.Id,
		RecipeId: rm.RecipeId,
		Name:     rm.Name,
	}
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
