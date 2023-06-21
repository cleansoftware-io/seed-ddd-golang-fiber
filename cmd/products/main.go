package products

import (
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/initialization"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/infra/adapters"
)
import "github.con/tgarcia/seed-ddd-golang-fiber/internal/infra/controllers"
import "github.con/tgarcia/seed-ddd-golang-fiber/internal/application"

func Run(bootstrap *initialization.Bootstrap) {

	logger := adapters.NewLoggerImpl(bootstrap.Logger)
	productsService := adapters.NewDeactivateProductsIml(logger)

	deactivateProductUseCase := application.NewDeactivateProductUseCase(logger, productsService)

	productsControllers := controllers.NewProducts(bootstrap.App, bootstrap.Logger, deactivateProductUseCase)
	productsControllers.RegisterRoutes()

}
