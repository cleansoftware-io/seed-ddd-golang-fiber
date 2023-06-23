//go:build wireinject
// +build wireinject

package main

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/initialization"
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/services"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
	"github.com/google/wire"
)

var productsServiceProviderSet wire.ProviderSet = wire.NewSet(
	adapters.ProvideDeactivateProductsIml,
	wire.Bind(new(services.Products), new(adapters.DeactivateProductsIml)))

func InitializeApplication() initialization.ProductInitialization {
	wire.Build(
		config.SetupFiber,
		config.SetupLogging,
		adapters.ProvideLoggerImpl,
		initialization.ProvideBootstrap,
		productsServiceProviderSet,
		initialization.ProvideProductInitialization)
	return initialization.ProductInitialization{}
}
