// recipe_material_model.go
package model

import (
	"time"

	"github.com/uptrace/bun"
)

// RecipeMaterial　ORMモデル
type RecipeMaterial struct {
	bun.BaseModel `bun:"table:recipe_materials"`

	Id        int64     `bun:"id,pk,type:uuid,default:gen_random_uuid()"`             // id
	RecipeId  int64     `bun:"recipe_id,notnull"`                                     // recipes.id
	Name      string    `bun:"name,notnull"`                                          // 材料名
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"` // 登録日時
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"` // 更新日時
}
