package initialization

import (
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/services"
)
import "cleansoftware.io/ddd/fiber/seed/internal/products/infra/controllers"
import "cleansoftware.io/ddd/fiber/seed/internal/products/application"

type ProductInitialization struct {
	Bootstrap Bootstrap
}

func ProvideProductInitialization(bootstrap Bootstrap, productsService services.Products) ProductInitialization {

	deactivateProductUseCase := application.NewDeactivateProductUseCase(bootstrap.Logger, productsService)

	// ... more use cases
	// ... more services
	// ... more controllers
	// This method is a little ugly, I prefer a DI container to manage all dependencies
	// to avoid complex constructors and complex app initialization.

	productsControllers := controllers.NewProducts(bootstrap.App, bootstrap.Logger, deactivateProductUseCase)
	productsControllers.RegisterRoutes()
	return ProductInitialization{
		Bootstrap: bootstrap,
	}
}
