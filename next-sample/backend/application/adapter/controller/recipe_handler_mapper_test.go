package controller

import (
	"reflect"
	"testing"
	"time"

	"github.com/react/next-sample/backend/infrastructure/openapi"
	"github.com/react/next-sample/backend/usecase/dto"
)

func TestToResponse(t *testing.T) {
	// Arrange
	now := time.Now()
	recipe := &dto.Recipe{
		Id:        1,
		Title:     "Test Recipe",
		Content:   "This is a test recipe content.",
		Image:     "test-image-url",
		CreatedAt: now,
		UpdatedAt: now,
		RecipeMaterials: []dto.RecipeMaterial{
			{Name: "Material 1"},
			{Name: "Material 2"},
		},
	}

	expected := &openapi.Recipe{
		Id:        1,
		Title:     "Test Recipe",
		Content:   "This is a test recipe content.",
		Image:     "test-image-url",
		CreatedAt: now,
		UpdatedAt: now,
		Materials: []string{"Material 1", "Material 2"},
	}

	// Act
	result := ToResponse(recipe)

	// Assert
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToResponse() = %v, want %v", result, expected)
	}
}
