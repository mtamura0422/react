package usecase

import (
	"testing"

	"github.com/react/next-sample/backend/domain/entity"
	repositories "github.com/react/next-sample/backend/domain/repositories/mock"
	"go.uber.org/mock/gomock"
)

func TestAddRecipe(t *testing.T) {

	create_data := &entity.Recipe{
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "tedt.png",
	}

	create_material_data := &entity.RecipeMaterial{
		RecipeId: 1,
		Name:     "材料",
	}

	ret_data := &entity.Recipe{
		Id:      1,
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "tedt.png",
	}

	ret_material_data := &entity.RecipeMaterial{
		Id:       1,
		RecipeId: 1,
		Name:     "材料",
	}

	// mockの作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeRepo := repositories.NewMockRecipeRepository(mockCtrl)
	mockRecipeRepo.EXPECT().Create(gomock.Any(), create_data).Return(
		ret_data, nil,
	)

	mockRecipeMateRepo := repositories.NewMockRecipeMaterialRepository(mockCtrl)
	mockRecipeMateRepo.EXPECT().Create(gomock.Any(), create_material_data).Return(
		ret_material_data, nil,
	)

	//var expect *entity.Recipe = ret_data
	/*
	   	usecase := &UserUsecase{
	   		repository: mock,
	   	}

	   // 結果チェック
	   result, err := usecase.repository.Create(user)

	   	if err != nil {
	   		t.Error("error happen")
	   	}

	   	if diff := cmp.Diff(result, expect); diff != "" {
	   		t.Errorf("User Data miss match :%s", diff)
	   	}
	*/
}
