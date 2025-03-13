package controller

import (
	"github.com/react/next-sample/backend/infrastructure/openapi"
	"github.com/react/next-sample/backend/usecase/dto"
)

func ToDTO(req *openapi.Recipe) *dto.Recipe {

	recipeMaterials := make([]dto.RecipeMaterial, len(req.Materials))

	for i, rm := range req.Materials {
		recipeMaterials[i] = dto.RecipeMaterial{
			Name: rm,
		}
	}

	return &dto.Recipe{
		Id:              req.Id,
		Title:           req.Title,
		Content:         req.Content,
		CreatedAt:       req.CreatedAt,
		UpdatedAt:       req.UpdatedAt,
		RecipeMaterials: recipeMaterials,
	}
}

func ToResponse(u *dto.Recipe) *openapi.Recipe {

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
