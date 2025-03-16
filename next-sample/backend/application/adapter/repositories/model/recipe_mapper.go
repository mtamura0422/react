package model

import (
	"fmt"
	"log"

	"github.com/react/next-sample/backend/domain/entity"
)

// ORM モデルを Entity に変換
func (r *Recipe) ToEntity() *entity.Recipe {
	log.Printf("recipe ToEntity tuuka1")
	rRecipeMaterial := make([]entity.RecipeMaterial, len(r.RecipeMaterials))
	log.Printf("recipe ToEntity tuuka2")
	for i, recipeMaterial := range r.RecipeMaterials {

		fmt.Printf("%T\n", recipeMaterial.ToEntity())
		rRecipeMaterial[i] = *recipeMaterial.ToEntity()

	}
	log.Printf("recipe ToEntity tuuka3")
	return &entity.Recipe{
		Id:        r.Id,
		Title:     r.Title,
		Filename:  r.Image,
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
		Image:     e.Filename,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
