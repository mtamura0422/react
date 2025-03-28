// recipe_material.go
package entity

import "time"

// RecipeMaterial eintityモデル
type RecipeMaterial struct {
	Id        int64
	RecipeId  int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
