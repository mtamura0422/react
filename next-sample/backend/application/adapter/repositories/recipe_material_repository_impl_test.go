// recipe_material_repository_impl_test.go
package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

/*
TestCreateMaterial レシピ登録テスト
*/
func TestCreateMaterial(t *testing.T) {

	t.Run("正常系 データあり", func(t *testing.T) {

		entity_data := &entity.RecipeMaterial{
			Id:        1,
			RecipeId:  1,
			Name:      "材料1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectExec(
			"INSERT INTO `recipe_materials`").
			WillReturnResult(sqlmock.NewResult(1, 1))

		rep := NewRecipeMaterialRepository(bun.NewDB(dbMock, mysqldialect.New()))

		// テスト実行
		ret, err := rep.Create(context.Background(), entity_data)

		// 戻り値の検証
		assert.NoError(t, err)
		assert.Equal(t, int64(1), ret.Id)
		assert.Equal(t, entity_data.Name, ret.Name)
		assert.Equal(t, entity_data.RecipeId, ret.RecipeId)

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestCreateRecipe: %v", err)
		}
	})
}
