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

	// ... more use cases
	// ... more services
	// ... more controllers
	// This method is a little ugly, I prefer a DI container to manage all dependencies
	// to avoid complex constructors and complex app initialization.

	productsControllers := controllers.NewProducts(bootstrap.App, bootstrap.Logger, deactivateProductUseCase)
	productsControllers.RegisterRoutes()

}
