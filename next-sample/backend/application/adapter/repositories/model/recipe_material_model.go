package model

import (
	"time"

	"github.com/uptrace/bun"
)

// ORM モデル
type RecipeMaterial struct {
	bun.BaseModel `bun:"table:recipe_materials"`

	Id        int64     `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	RecipeId  int64     `bun:"recipe_id,notnull"`
	Name      string    `bun:"name,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
