package products

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/initialization"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
)
import "cleansoftware.io/ddd/fiber/seed/internal/products/infra/controllers"
import "cleansoftware.io/ddd/fiber/seed/internal/products/application"

func Run(bootstrap initialization.Bootstrap) {

	logger := bootstrap.Logger

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
