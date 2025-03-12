package service

import (
	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/infrastructure/openapi"
)

func ToEntity(r *openapi.Recipe) *entity.Recipe {

	return &entity.Recipe{
		Id:      r.Id,
		Title:   r.Title,
		Content: r.Content,
	}
}

func ToEntityMaterial(r *openapi.Recipe) []*entity.RecipeMaterial {

	recipeMaterial := make([]*entity.RecipeMaterial, len(r.Materials))

	for i, rm := range r.Materials {
		recipeMaterial[i] = &entity.RecipeMaterial{
			Name: rm,
		}
	}

	return recipeMaterial

}

func ToResponse(e *entity.Recipe) *openapi.Recipe {

	materials := make([]string, len(e.RecipeMaterials))

	for i, rm := range e.RecipeMaterials {
		materials[i] = rm.Name
	}

	return &openapi.Recipe{
		Id:        e.Id,
		Title:     e.Title,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Materials: materials,
	}
}
