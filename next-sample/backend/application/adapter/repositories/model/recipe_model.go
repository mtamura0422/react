// model package
// recipe_model.go
package model

import (
	"time"

	"github.com/uptrace/bun"
)

// Recipe　ORMモデル
type Recipe struct {
	bun.BaseModel `bun:"table:recipes"`

	Id        int64     `bun:"id,pk,type:uuid,default:gen_random_uuid()"`             // id
	Title     string    `bun:"title,notnull"`                                         // タイトル
	Content   string    `bun:"content,notnull"`                                       // 作り方
	Image     string    `bun:"image"`                                                 // 画像パス
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"` // 登録日時
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"` // 更新日時

	RecipeMaterials []RecipeMaterial `bun:"rel:has-many,join:id=recipe_id"` // 材料
}
