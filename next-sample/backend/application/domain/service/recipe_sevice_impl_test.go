package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/react/next-sample/backend/domain/entity"
	repositories "github.com/react/next-sample/backend/domain/repositories/mock"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateRecipe(t *testing.T) {
	// ★ モックデータの準備
	create_material_data := &entity.RecipeMaterial{
		RecipeId: 1,
		Name:     "材料",
	}

	entityRecipeMaterial := []entity.RecipeMaterial{*create_material_data}

	create_data := &entity.Recipe{
		Title:           "テストタイトル",
		Content:         "テスト作り方",
		Image:           "test.png",
		RecipeMaterials: entityRecipeMaterial,
	}

	ret_data := &entity.Recipe{
		Id:      1,
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "test.png",
	}

	ret_material_data := &entity.RecipeMaterial{
		Id:       1,
		RecipeId: 1,
		Name:     "材料",
	}

	retRecipeMaterial := []entity.RecipeMaterial{*ret_material_data}

	ret_create_data := ret_data
	ret_create_data.RecipeMaterials = retRecipeMaterial

	// ★ モックの作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeMateRepo := repositories.NewMockRecipeMaterialRepository(mockCtrl)

	// ★ モックの期待する呼び出しを定義
	mockRecipeRepo.EXPECT().
		Create(gomock.Any(), gomock.Eq(create_data)).
		Return(ret_data, nil).
		Times(1) // 1回だけ呼び出し

	mockRecipeMateRepo.EXPECT().
		Create(gomock.Any(), gomock.Eq(create_material_data)).
		Return(ret_material_data, nil).
		Times(1) // 1回だけ呼び出し

	// ★ モックDBとTxRepositoryを生成
	bunDB, err := db.NewDBMock()
	assert.NoError(t, err)

	txRepository := db.NewTxRepository(bunDB)

	// ★ サービス呼び出し
	mockService := NewRecipeService(txRepository, mockRecipeRepo, mockRecipeMateRepo)

	// ★ モックを通してCreateRecipeTxを実行 → モックが呼び出されるか検証
	result, err := mockService.CreateRecipeTx(context.Background(), create_data)
	assert.NoError(t, err)

	// ★ 返り値の検証
	assert.NotNil(t, result)

	if !reflect.DeepEqual(result, ret_create_data) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", ret_create_data, result)
	}

}
