package dto

import (
	"reflect"
	"testing"

	"github.com/react/next-sample/backend/domain/entity"
)

// DTO モデルを Entity に変換
func TestToEntity(t *testing.T) {

	dto_data := &Recipe{
		Id:      1,
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "test.png",
		RecipeMaterials: []RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料1",
			},
		},
	}

	RecipeImage := entity.RecipeImage{
		File:       dto_data.RecipeImage.File,
		FileHeader: nil,
	}

	check_data := &entity.Recipe{
		Id:      dto_data.Id,
		Title:   dto_data.Title,
		Content: dto_data.Content,
		RecipeMaterials: []entity.RecipeMaterial{
			{

				Name: "材料1",
			},
		},
		RecipeImage: RecipeImage,
	}

	entity_data := dto_data.ToEntity()

	if !reflect.DeepEqual(check_data, entity_data) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", check_data, entity_data)
	}
}

// Entity を DTO(DataTransferObject)に変換
func TestToRecipeMapper(t *testing.T) {

	entity_data := &entity.Recipe{
		Id:      1,
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "test.png",
		RecipeMaterials: []entity.RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料1",
			},
		},
	}

	filename := IMAGE_DMAIN + IMAGE_PATH
	filename += entity_data.Image

	check_data := &Recipe{
		Id:        entity_data.Id,
		Title:     entity_data.Title,
		Content:   entity_data.Content,
		Image:     filename,
		CreatedAt: entity_data.CreatedAt,
		UpdatedAt: entity_data.UpdatedAt,
		RecipeMaterials: []RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料1",
			},
		},
	}

	ret_data := ToRecipeMapper(entity_data)

	if !reflect.DeepEqual(check_data, ret_data) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", check_data, ret_data)
	}
}
