//go:build wireinject
// +build wireinject

package main

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/initialization"
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/services"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
	"github.com/google/wire"
)

var productsServiceProviderSet wire.ProviderSet = wire.NewSet(
	adapters.ProvideDeactivateProductsIml,
	wire.Bind(new(ports.Products), new(adapters.DeactivateProductsIml)))

func InitializeApplication() initialization.ProductInitialization {
	wire.Build(
		config.SetupFiber,
		config.SetupLogging,
		adapters.ProvideLoggerImpl,
		initialization.ProvideBootstrap,
		productsServiceProviderSet,
		services.ProvideServiceList,
		initialization.ProvideProductInitialization)
	return initialization.ProductInitialization{}
}
