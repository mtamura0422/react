package entity

import (
	"mime/multipart"
	"time"
)

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

type RecipeImage struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}
