package repositories

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func TestCreate(t *testing.T) {

	entity_data := &entity.Recipe{
		Title:   "テストタイトル",
		Content: "テスト作り方",
		Image:   "test.png",
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
		"INSERT INTO `recipes`").
		WillReturnResult(sqlmock.NewResult(1, 1))

	rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))

	// テスト実行
	ret, err := rep.Create(context.Background(), entity_data)

	// 戻り値の検証
	assert.NoError(t, err)
	assert.Equal(t, int64(1), ret.Id)
	assert.Equal(t, entity_data.Title, ret.Title)
	assert.Equal(t, entity_data.Content, ret.Content)
	assert.Equal(t, entity_data.Image, ret.Image)

	// モックの期待通りに呼び出されたか検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("TestCreateRecipe: %v", err)
	}
}

func TestGet(t *testing.T) {

	t.Run("正常系 データあり", func(t *testing.T) {
		recipe_data := &entity.Recipe{
			Id:        1,
			Title:     "テストタイトル",
			Content:   "テスト作り方",
			Image:     "test.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		material_data := &entity.RecipeMaterial{
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

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` WHERE (id = 1)")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "image", "created_at", "updated_at"}).
				AddRow(recipe_data.Id, recipe_data.Title, recipe_data.Content, recipe_data.Image, recipe_data.CreatedAt, recipe_data.UpdatedAt))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe_material`.`id`, `recipe_material`.`recipe_id`, `recipe_material`.`name`, `recipe_material`.`created_at`, `recipe_material`.`updated_at` FROM `recipe_materials` AS `recipe_material` WHERE (`recipe_material`.`recipe_id` IN (1))")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "recipe_id", "name", "created_at", "updated_at"}).
				AddRow(material_data.Id, material_data.RecipeId, material_data.Name, material_data.CreatedAt, material_data.UpdatedAt))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, err := rep.Get(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, recipe_data.Id, ret.Id)
		assert.Equal(t, recipe_data.Title, ret.Title)
		assert.Equal(t, recipe_data.Content, ret.Content)
		assert.Equal(t, recipe_data.Image, ret.Image)
		assert.Equal(t, recipe_data.CreatedAt, ret.CreatedAt)
		assert.Equal(t, recipe_data.UpdatedAt, ret.UpdatedAt)

		assert.Equal(t, material_data.Id, ret.RecipeMaterials[0].Id)
		assert.Equal(t, material_data.RecipeId, ret.RecipeMaterials[0].RecipeId)
		assert.Equal(t, material_data.Name, ret.RecipeMaterials[0].Name)
		assert.Equal(t, material_data.CreatedAt, ret.RecipeMaterials[0].CreatedAt)
		assert.Equal(t, material_data.UpdatedAt, ret.RecipeMaterials[0].UpdatedAt)

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestGetRecipe: %v", err)
		}
	})

	t.Run("異常系 データなし", func(t *testing.T) {

		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` WHERE (id = 20)"))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, err := rep.Get(context.Background(), 20)

		assert.Error(t, err)
		assert.Nil(t, ret)

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestGetRecipe: %v", err)
		}
	})

}

func TestList(t *testing.T) {

	t.Run("正常系 データあり", func(t *testing.T) {
		entity_data := []*entity.Recipe{
			{
				Id:        1,
				Title:     "テストタイトル",
				Content:   "テスト作り方",
				Image:     "test.png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        2,
				Title:     "テストタイトル",
				Content:   "テスト作り方",
				Image:     "test.png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		material_data := []*entity.RecipeMaterial{
			{
				Id:        1,
				RecipeId:  1,
				Name:      "材料1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        2,
				RecipeId:  2,
				Name:      "材料2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `recipes` AS `recipe`")).
			WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(2))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` ORDER BY `id` desc LIMIT 10")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "image", "created_at", "updated_at"}).
				AddRow(entity_data[0].Id, entity_data[0].Title, entity_data[0].Content, entity_data[0].Image, entity_data[0].CreatedAt, entity_data[0].UpdatedAt).
				AddRow(entity_data[1].Id, entity_data[1].Title, entity_data[1].Content, entity_data[1].Image, entity_data[1].CreatedAt, entity_data[1].UpdatedAt))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe_material`.`id`, `recipe_material`.`recipe_id`, `recipe_material`.`name`, `recipe_material`.`created_at`, `recipe_material`.`updated_at` FROM `recipe_materials` AS `recipe_material` WHERE (`recipe_material`.`recipe_id` IN (1, 2))")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "recipe_id", "name", "created_at", "updated_at"}).
				AddRow(material_data[0].Id, material_data[0].RecipeId, material_data[0].Name, material_data[0].CreatedAt, material_data[0].UpdatedAt).
				AddRow(material_data[1].Id, material_data[1].RecipeId, material_data[1].Name, material_data[1].CreatedAt, material_data[1].UpdatedAt))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, ret_num, err := rep.List(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, ret)

		require.Len(t, ret, 2)
		require.Equal(t, 2, ret_num)

		for i, p := range ret {

			require.Equal(t, 1, len(p.RecipeMaterials))

			assert.Equal(t, entity_data[i].Id, p.Id)
			assert.Equal(t, entity_data[i].Title, p.Title)
			assert.Equal(t, entity_data[i].Content, p.Content)
			assert.Equal(t, entity_data[i].Image, p.Image)
			assert.Equal(t, entity_data[i].CreatedAt, p.CreatedAt)
			assert.Equal(t, entity_data[i].UpdatedAt, p.UpdatedAt)

			assert.Equal(t, material_data[i].Id, p.RecipeMaterials[0].Id)
			assert.Equal(t, material_data[i].RecipeId, p.RecipeMaterials[0].RecipeId)
			assert.Equal(t, material_data[i].Name, p.RecipeMaterials[0].Name)
			assert.Equal(t, material_data[i].CreatedAt, p.RecipeMaterials[0].CreatedAt)
			assert.Equal(t, material_data[i].UpdatedAt, p.RecipeMaterials[0].UpdatedAt)
		}

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestListRecipe: %v", err)
		}
	})

	t.Run("異常系 データなし", func(t *testing.T) {
		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `recipes` AS `recipe`")).
			WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(0))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` ORDER BY `id` desc LIMIT 10"))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, ret_num, err := rep.List(context.Background(), 1)

		assert.Error(t, err)
		assert.Nil(t, ret)

		require.Len(t, ret, 0)
		require.Equal(t, 0, ret_num)
	})
}

func TestSearch(t *testing.T) {

	t.Run("正常系 データあり", func(t *testing.T) {
		entity_data := []*entity.Recipe{
			{
				Id:        1,
				Title:     "テストタイトル",
				Content:   "テスト作り方",
				Image:     "test.png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        2,
				Title:     "テストタイトル",
				Content:   "テスト作り方",
				Image:     "test.png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		material_data := []*entity.RecipeMaterial{
			{
				Id:        1,
				RecipeId:  1,
				Name:      "材料1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        2,
				RecipeId:  2,
				Name:      "材料2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `recipes` AS `recipe` WHERE ((title LIKE '%テスト%') or (content LIKE '%テスト%'))")).
			WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(2))

		mock.ExpectQuery(regexp.QuoteMeta("recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` WHERE ((title LIKE '%テスト%') or (content LIKE '%テスト%')) ORDER BY `id` desc LIMIT 10")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "image", "created_at", "updated_at"}).
				AddRow(entity_data[0].Id, entity_data[0].Title, entity_data[0].Content, entity_data[0].Image, entity_data[0].CreatedAt, entity_data[0].UpdatedAt).
				AddRow(entity_data[1].Id, entity_data[1].Title, entity_data[1].Content, entity_data[1].Image, entity_data[1].CreatedAt, entity_data[1].UpdatedAt))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT `recipe_material`.`id`, `recipe_material`.`recipe_id`, `recipe_material`.`name`, `recipe_material`.`created_at`, `recipe_material`.`updated_at` FROM `recipe_materials` AS `recipe_material` WHERE (`recipe_material`.`recipe_id` IN (1, 2))")).
			WillReturnRows(sqlmock.NewRows([]string{"id", "recipe_id", "name", "created_at", "updated_at"}).
				AddRow(material_data[0].Id, material_data[0].RecipeId, material_data[0].Name, material_data[0].CreatedAt, material_data[0].UpdatedAt).
				AddRow(material_data[1].Id, material_data[1].RecipeId, material_data[1].Name, material_data[1].CreatedAt, material_data[1].UpdatedAt))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, ret_num, err := rep.Search(context.Background(), "テスト", 1)

		assert.NoError(t, err)
		assert.NotNil(t, ret)

		require.Len(t, ret, 2)
		require.Equal(t, 2, ret_num)

		for i, p := range ret {

			require.Equal(t, 1, len(p.RecipeMaterials))

			assert.Equal(t, entity_data[i].Id, p.Id)
			assert.Equal(t, entity_data[i].Title, p.Title)
			assert.Equal(t, entity_data[i].Content, p.Content)
			assert.Equal(t, entity_data[i].Image, p.Image)
			assert.Equal(t, entity_data[i].CreatedAt, p.CreatedAt)
			assert.Equal(t, entity_data[i].UpdatedAt, p.UpdatedAt)

			assert.Equal(t, material_data[i].Id, p.RecipeMaterials[0].Id)
			assert.Equal(t, material_data[i].RecipeId, p.RecipeMaterials[0].RecipeId)
			assert.Equal(t, material_data[i].Name, p.RecipeMaterials[0].Name)
			assert.Equal(t, material_data[i].CreatedAt, p.RecipeMaterials[0].CreatedAt)
			assert.Equal(t, material_data[i].UpdatedAt, p.RecipeMaterials[0].UpdatedAt)
		}

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestSearchRecipe: %v", err)
		}
	})

	t.Run("異常系 データなし", func(t *testing.T) {
		// モックDBを作成 (*sql.DBを返す)
		dbMock, mock, err := db.NewMockDB()
		if err != nil {
			t.Fatalf("failed to create mock DB: %v", err)
		}
		defer dbMock.Close()

		mock.ExpectQuery("SELECT version()").
			WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.33"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `recipes` AS `recipe` WHERE ((title LIKE '%テスト%') or (content LIKE '%テスト%'))")).
			WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(0))

		mock.ExpectQuery(regexp.QuoteMeta("recipe`.`id`, `recipe`.`title`, `recipe`.`content`, `recipe`.`image`, `recipe`.`created_at`, `recipe`.`updated_at` FROM `recipes` AS `recipe` WHERE ((title LIKE '%テスト%') or (content LIKE '%テスト%')) ORDER BY `id` desc LIMIT 10"))

		rep := NewRecipeRepository(bun.NewDB(dbMock, mysqldialect.New()))
		ret, ret_num, err := rep.Search(context.Background(), "テスト", 1)
		log.Printf("aaaaa%v", ret)
		assert.Error(t, err)
		assert.Nil(t, ret)

		require.Len(t, ret, 0)
		require.Equal(t, 0, ret_num)

		// モックの期待通りに呼び出されたか検証
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("TestSearchRecipe: %v", err)
		}
	})
}
