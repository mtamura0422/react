package entity

import "time"

type Recipe struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time

	RecipeMaterials []RecipeMaterial
}
