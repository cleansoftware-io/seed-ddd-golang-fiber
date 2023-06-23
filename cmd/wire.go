//go:build wireinject
// +build wireinject

package main

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/initialization"
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"cleansoftware.io/ddd/fiber/seed/internal/products/infra/adapters"
	"github.com/google/wire"
)

func InitializeApplication() initialization.Bootstrap {
	wire.Build(config.SetupFiber, config.SetupLogging, adapters.ProvideLoggerImpl, initialization.ProvideBootstrap)
	return initialization.Bootstrap{}
}
