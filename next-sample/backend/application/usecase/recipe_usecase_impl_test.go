package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/react/next-sample/backend/domain/entity"
	repositories "github.com/react/next-sample/backend/domain/repositories/mock"
	service "github.com/react/next-sample/backend/domain/service/port/mock"
	"github.com/react/next-sample/backend/usecase/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddRecipe(t *testing.T) {
	// ★ モックデータの準備
	create_material_data := &dto.RecipeMaterial{
		RecipeId: 1,
		Name:     "材料",
	}

	dtoRecipeMaterial := []dto.RecipeMaterial{*create_material_data}

	create_data := &dto.Recipe{
		Title:           "テストタイトル",
		Content:         "テスト作り方",
		Image:           "test.png",
		RecipeMaterials: dtoRecipeMaterial,
	}

	ret_entity_data := &entity.Recipe{
		Id:      1,
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "test.png", // ✅ モックの期待値と一致
		RecipeMaterials: []entity.RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料",
			},
		},
	}

	// ★ モック作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeSearvice := service.NewMockRecipeService(mockCtrl)

	// ★ DTO → Entity の変換
	entityData := create_data.ToEntity()

	// ★ モックの期待する呼び出しを定義
	mockRecipeSearvice.EXPECT().
		CreateRecipeTx(gomock.Any(), gomock.Eq(entityData)).
		Return(ret_entity_data, nil).
		MinTimes(1)

	mockUsecase := NewRecipeUsecase(mockRecipeRepo, mockRecipeSearvice)

	// ★ テスト実行
	result, err := mockUsecase.AddRecipe(context.Background(), create_data)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	if !reflect.DeepEqual(result.Title, create_data.Title) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", create_data, result)
	}
}
