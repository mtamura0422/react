//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/react/next-sample/backend/adapter/controller"
	"github.com/react/next-sample/backend/adapter/repositories"
	"github.com/react/next-sample/backend/domain/service"
	"github.com/react/next-sample/backend/infrastructure/db"
	"github.com/react/next-sample/backend/usecase"
)

var BunDBSet = wire.NewSet(
	db.NewDB,
)

func InitializeServer() *controller.Server {

	wire.Build(
		BunDBSet,
		db.NewTxRepository,
		repositories.NewRecipeRepository,
		repositories.NewRecipeMaterialRepository,
		service.NewRecipeService,
		usecase.NewRecipeUsecase,
		usecase.NewRecipeMaterialUsecase,
		controller.NewServer,
	)
	return &controller.Server{}
}
