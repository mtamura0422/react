package controller

import (
	"github.com/react/next-sample/backend/infrastructure/openapi"
	"github.com/react/next-sample/backend/usecase"
)

func ToDTO(req *openapi.Recipe) *usecase.Recipe {
	return &usecase.Recipe{
		Id:        req.Id,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}

func ToResponse(u *usecase.Recipe) *openapi.Recipe {

	materials := make([]string, len(u.RecipeMaterials))

	for i, rm := range u.RecipeMaterials {
		materials[i] = rm.Name
	}

	return &openapi.Recipe{
		Id:        u.Id,
		Title:     u.Title,
		Content:   u.Content,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Materials: materials,
	}
}
