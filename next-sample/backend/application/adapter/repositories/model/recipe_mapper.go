package model

import (
	"fmt"

	"github.com/react/next-sample/backend/domain/entity"
)

// ORM モデルを Entity に変換
func (r *Recipe) ToEntity() *entity.Recipe {

	rRecipeMaterial := make([]entity.RecipeMaterial, len(r.RecipeMaterials))

	for i, recipeMaterial := range r.RecipeMaterials {

		fmt.Printf("%T\n", recipeMaterial.ToEntity())
		rRecipeMaterial[i] = *recipeMaterial.ToEntity()

	}

	return &entity.Recipe{
		Id:        r.Id,
		Title:     r.Title,
		Content:   r.Content,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,

		RecipeMaterials: rRecipeMaterial,
	}
}

// Entity を ORM モデルに変換
func ToRecipeMapper(e *entity.Recipe) *Recipe {
	return &Recipe{
		Id:        e.Id,
		Title:     e.Title,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
