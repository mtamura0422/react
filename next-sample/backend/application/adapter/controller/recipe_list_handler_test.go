package controller

import (
	"encoding/json"
	"fmt"
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

func TestGetRecipeList(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define mock parameters and response
	params := openapi.GetRecipeListParams{
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
	mockRecipeUsecase.EXPECT().GetRecipeList(gomock.Any(), int64(*params.Page)).Return(mockRecipes, totalCount, nil)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/list", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipeList handler
	server.GetRecipeList(rr, req, params)

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

func TestGetRecipeList_Error(t *testing.T) {
	// Create a new Server instance with mocked dependencies
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRecipeUsecase := port.NewMockRecipeUsecase(mockCtrl)
	server := &Server{
		recipeUsecase: mockRecipeUsecase,
	}

	// Define mock parameters
	params := openapi.GetRecipeListParams{
		Page: func(i int) *int { return &i }(1),
	}

	// Set up the mock behavior to return an error
	mockRecipeUsecase.EXPECT().GetRecipeList(gomock.Any(), int64(*params.Page)).Return(nil, 0, fmt.Errorf("search error"))

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/recipes/list", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetRecipeList handler
	server.GetRecipeList(rr, req, params)

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
