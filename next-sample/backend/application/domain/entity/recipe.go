// recipe.go
package entity

import (
	"mime/multipart"
	"time"
)

// Recipe レシピentityモデル
type Recipe struct {
	Id        int64
	Title     string
	Content   string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time

	RecipeMaterials []RecipeMaterial
	RecipeImage     RecipeImage
}

// RecipeImage レシピ画像entityモデル
type RecipeImage struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}
