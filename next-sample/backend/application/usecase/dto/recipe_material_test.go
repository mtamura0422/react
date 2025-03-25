package dto

import (
	"reflect"
	"testing"

	"github.com/react/next-sample/backend/domain/entity"
)

// Entity を ORM モデルに変換
func TestToRecipeMaterialMapper(t *testing.T) {

	entity_data := &entity.RecipeMaterial{
		Id:       1,
		RecipeId: 1,
		Name:     "材料1",
	}

	ret_data := ToRecipeMaterialMapper(entity_data)

	check_data := &RecipeMaterial{
		Id:        entity_data.Id,
		RecipeId:  entity_data.RecipeId,
		Name:      entity_data.Name,
		CreatedAt: entity_data.CreatedAt,
		UpdatedAt: entity_data.UpdatedAt,
	}

	if !reflect.DeepEqual(check_data, ret_data) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", check_data, ret_data)
	}
}
