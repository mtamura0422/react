package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/react/next-sample/backend/infrastructure/openapi"
	"github.com/react/next-sample/backend/usecase/dto"
	port "github.com/react/next-sample/backend/usecase/port/mock"
	"go.uber.org/mock/gomock"
)

func TestGetVersion(t *testing.T) {
	// Create a new Server instance
	server := &Server{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetVersion handler
	server.GetVersion(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("RegisterRecipe handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "GetVersion"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRegisterRecipe(t *testing.T) {
	// モックの作成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// モックの期待値を設定
	mockRecipe := &dto.Recipe{
		Id:        1,
		Title:     "Test Recipe",
		Content:   "Test instructions",
		Image:     "testfile.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		RecipeMaterials: []dto.RecipeMaterial{
			{Name: "Ingredient1"},
			{Name: "Ingredient2"},
		},
	}
	mockRecipeUsecase.EXPECT().AddRecipe(gomock.Any(), gomock.Any()).Return(mockRecipe, nil)

	// マルチパートフォームデータを作成
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	writer.WriteField("title", "Test Recipe")

	materials := []string{"Ingredient1", "Ingredient2"}
	materialsJSON, err := json.Marshal(materials)
	if err != nil {
		t.Fatalf("Failed to marshal materials: %v", err)
	}
	writer.WriteField("materials", string(materialsJSON))
	writer.WriteField("content", "Test instructions")

	mockFileContent := []byte("mock file content")
	part, err := writer.CreateFormFile("file", "test.png")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	if _, err := part.Write(mockFileContent); err != nil {
		t.Fatalf("Failed to write mock file content: %v", err)
	}
	writer.Close()

	// HTTPリクエストを作成
	req, err := http.NewRequest(http.MethodPost, "/recipes", &body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// レスポンスを記録
	rr := httptest.NewRecorder()

	// ハンドラーを呼び出し
	server.RegisterRecipe(rr, req)

	// ステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedResponse := map[string]interface{}{
		"content":    mockRecipe.Content,
		"id":         mockRecipe.Id,
		"image":      mockRecipe.Image,
		"created_at": mockRecipe.CreatedAt,
		"updated_at": mockRecipe.UpdatedAt,
		"materials":  []string{"Ingredient1", "Ingredient2"},
		"title":      mockRecipe.Title,
	}
	expected, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Failed to marshal expected response: %v", err)
	}

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}

}
func TestGetRecipe(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define a valid recipe ID and mock response
	recipeID := int64(1)
	mockRecipe := &dto.Recipe{
		Id:        recipeID,
		Title:     "Test Recipe",
		Content:   "テスト作り方",
		Image:     "test1.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		RecipeMaterials: []dto.RecipeMaterial{
			{
				Id:       1,
				RecipeId: 1,
				Name:     "材料1",
			},
		},
	}

	// Set up the mock behavior
	mockRecipeUsecase.EXPECT().FindRecipe(context.Background(), recipeID).Return(mockRecipe, nil)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipe handler
	server.GetRecipe(rr, req, recipeID)

	// Check the response body
	expectedResponse := map[string]interface{}{
		"content":    mockRecipe.Content,
		"id":         mockRecipe.Id,
		"image":      mockRecipe.Image,
		"created_at": mockRecipe.CreatedAt,
		"updated_at": mockRecipe.UpdatedAt,
		"materials":  []string{"材料1"},
		"title":      mockRecipe.Title,
	}
	expected, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Failed to marshal expected response: %v", err)
	}

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}

}

func TestGetRecipe_NotFound(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define a recipe ID that does not exist
	recipeID := int64(999)

	// Set up the mock behavior to return an error
	mockRecipeUsecase.EXPECT().FindRecipe(gomock.Any(), recipeID).Return(nil, fmt.Errorf("recipe not found"))

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/999", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipe handler
	server.GetRecipe(rr, req, recipeID)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body for an error message
	expected := `{"code":"pkg.CodeInternalServerError","error":"recipe not found"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestGetRecipeSearch(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define mock parameters and response
	params := openapi.GetRecipeSearchParams{
		Q:    "test",
		Page: func(i int) *int { return &i }(1),
	}
	mockRecipes := []*dto.Recipe{
		{
			Id:        1,
			Title:     "Test Recipe 1",
			Content:   "Test Content 1",
			Image:     "test1.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			RecipeMaterials: []dto.RecipeMaterial{
				{
					Id:       1,
					RecipeId: 1,
					Name:     "材料1",
				},
			},
		},
		{
			Id:        2,
			Title:     "Test Recipe 2",
			Content:   "Test Content 2",
			Image:     "test2.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			RecipeMaterials: []dto.RecipeMaterial{
				{
					Id:       1,
					RecipeId: 1,
					Name:     "材料2",
				},
			},
		},
	}
	totalCount := 2

	// Set up the mock behavior
	mockRecipeUsecase.EXPECT().SearchRecipeList(gomock.Any(), params.Q, int64(*params.Page)).Return(mockRecipes, totalCount, nil)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/search", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipeSearch handler
	server.GetRecipeSearch(rr, req, params)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedResponse := openapi.RecipeList{
		List: []openapi.Recipe{
			{
				Id:        1,
				Title:     mockRecipes[0].Title,
				Content:   mockRecipes[0].Content,
				Image:     mockRecipes[0].Image,
				CreatedAt: mockRecipes[0].CreatedAt,
				UpdatedAt: mockRecipes[0].UpdatedAt,
				Materials: []string{"材料1"},
			},
			{
				Id:        2,
				Title:     mockRecipes[1].Title,
				Content:   mockRecipes[1].Content,
				Image:     mockRecipes[1].Image,
				CreatedAt: mockRecipes[1].CreatedAt,
				UpdatedAt: mockRecipes[1].UpdatedAt,
				Materials: []string{"材料2"},
			},
		},
		Page:       1,
		PerPage:    10,
		TotalCount: totalCount,
	}
	expected, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Failed to marshal expected response: %v", err)
	}
	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestGetRecipeSearch_Error(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define mock parameters
	params := openapi.GetRecipeSearchParams{
		Q:    "test",
		Page: func(i int) *int { return &i }(1),
	}

	// Set up the mock behavior to return an error
	mockRecipeUsecase.EXPECT().SearchRecipeList(gomock.Any(), params.Q, int64(*params.Page)).Return(nil, 0, fmt.Errorf("search error"))

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/search", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipeSearch handler
	server.GetRecipeSearch(rr, req, params)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body for an error message
	expected := `{"code":"pkg.CodeInternalServerError","error":"search error"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
