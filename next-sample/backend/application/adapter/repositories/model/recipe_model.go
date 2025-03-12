package model

import (
	"time"

	"github.com/uptrace/bun"
)

// ORM モデル
type Recipe struct {
	bun.BaseModel `bun:"table:recipes"`

	Id        int64     `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Title     string    `bun:"title,notnull"`
	Content   string    `bun:"content,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`

	RecipeMaterials []RecipeMaterial `bun:"rel:has-many,join:id=recipe_id"`
}
