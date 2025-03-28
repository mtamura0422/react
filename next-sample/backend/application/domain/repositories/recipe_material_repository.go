// recipe_material_repository.go
package repositories

import (
	"context"

	"github.com/react/next-sample/backend/domain/entity"
)

//go:generate mockgen -source=./recipe_material_repository.go -destination=./mock/recipe_material_repository_mock.go -package=repositories

// RecipeMaterialRepository intereface
type RecipeMaterialRepository interface {
	Get(ctx context.Context, id int64) (*entity.RecipeMaterial, error)
	Create(ctx context.Context, recipe *entity.RecipeMaterial) (*entity.RecipeMaterial, error)
}
