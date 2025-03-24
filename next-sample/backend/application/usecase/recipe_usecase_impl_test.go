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
	"github.com/stretchr/testify/require"
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
		Image:   "test.png", // モックの期待値と一致
		RecipeMaterials: []entity.RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料",
			},
		},
	}

	// モック作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeSearvice := service.NewMockRecipeService(mockCtrl)

	// DTO → Entity の変換
	entityData := create_data.ToEntity()

	// モックの期待する呼び出しを定義
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

func TestFindRecipe(t *testing.T) {

	// ★ モックデータの準備
	create_material_data := &dto.RecipeMaterial{
		RecipeId: 1,
		Name:     "材料",
	}

	dtoRecipeMaterial := []dto.RecipeMaterial{*create_material_data}

	ret_data := &dto.Recipe{
		Title:           "テストタイトル",
		Content:         "テスト作り方",
		Image:           "test.png",
		RecipeMaterials: dtoRecipeMaterial,
	}

	entityData := ret_data.ToEntity()

	// モック作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeSearvice := service.NewMockRecipeService(mockCtrl)

	mockUsecase := NewRecipeUsecase(mockRecipeRepo, mockRecipeSearvice)

	// モックの期待する呼び出しを定義
	mockRecipeRepo.EXPECT().
		Get(gomock.Any(), int64(1)).
		Return(entityData, nil).
		MinTimes(1)

	// ★ テスト実行
	result, err := mockUsecase.FindRecipe(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	if !reflect.DeepEqual(result.Title, ret_data.Title) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", ret_data, result)
	}
}

func TestGetRecipeList(t *testing.T) {

	entityData := []*entity.Recipe{
		{
			Id:      1,
			Title:   "test-title-1",
			Content: "作り方1",
			Image:   "image1.png",
			RecipeMaterials: []entity.RecipeMaterial{
				{
					Id:       1,
					RecipeId: 1,
					Name:     "材料1",
				},
				{
					Id:       2,
					RecipeId: 1,
					Name:     "材料2",
				},
			},
		},
		{
			Id:      2,
			Title:   "test-title-2",
			Content: "作り方2",
			Image:   "image2.png",
			RecipeMaterials: []entity.RecipeMaterial{
				{
					Id:       3,
					RecipeId: 2,
					Name:     "材料3",
				},
			},
		},
	}

	// モック作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeSearvice := service.NewMockRecipeService(mockCtrl)

	mockUsecase := NewRecipeUsecase(mockRecipeRepo, mockRecipeSearvice)

	// モックの期待する呼び出しを定義
	mockRecipeRepo.EXPECT().
		List(gomock.Any(), int64(1)).
		Return(entityData, 2, nil).
		MinTimes(1)

	// ★ テスト実行
	result, ret_num, err := mockUsecase.GetRecipeList(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	require.Len(t, result, 2)
	require.Equal(t, 2, ret_num)

	for i, p := range result {
		require.Equal(t, len(entityData[i].RecipeMaterials), len(p.RecipeMaterials))

		checkData := dto.ToRecipeMapper(entityData[i])
		if !reflect.DeepEqual(p, checkData) {
			t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", checkData, p)
		}
	}
}
func TestSearchRecipeList(t *testing.T) {

	entityData := []*entity.Recipe{
		{
			Id:      1,
			Title:   "test-title-1",
			Content: "作り方1",
			Image:   "image1.png",
			RecipeMaterials: []entity.RecipeMaterial{
				{
					Id:       1,
					RecipeId: 1,
					Name:     "材料1",
				},
				{
					Id:       2,
					RecipeId: 1,
					Name:     "材料2",
				},
			},
		},
		{
			Id:      2,
			Title:   "test-title-2",
			Content: "作り方2",
			Image:   "image2.png",
			RecipeMaterials: []entity.RecipeMaterial{
				{
					Id:       3,
					RecipeId: 2,
					Name:     "材料3",
				},
			},
		},
	}

	// モック作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeSearvice := service.NewMockRecipeService(mockCtrl)

	mockUsecase := NewRecipeUsecase(mockRecipeRepo, mockRecipeSearvice)

	// モックの期待する呼び出しを定義
	mockRecipeRepo.EXPECT().
		Search(gomock.Any(), "test", int64(1)).
		Return(entityData, 2, nil).
		MinTimes(1)

	// ★ テスト実行
	result, ret_num, err := mockUsecase.SearchRecipeList(context.Background(), "test", 1)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	require.Len(t, result, 2)
	require.Equal(t, 2, ret_num)

	for i, p := range result {
		require.Equal(t, len(entityData[i].RecipeMaterials), len(p.RecipeMaterials))

		checkData := dto.ToRecipeMapper(entityData[i])
		if !reflect.DeepEqual(p, checkData) {
			t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", checkData, p)
		}
	}
}
