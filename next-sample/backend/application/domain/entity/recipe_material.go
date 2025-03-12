package entity

import "time"

type RecipeMaterial struct {
	Id        int64
	RecipeId  int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
