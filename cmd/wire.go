//go:build wireinject
// +build wireinject

package main

import (
	"cleansoftware.io/ddd/fiber/seed/cmd/initialization"
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"github.com/google/wire"
)

func InitializeApplication() initialization.Bootstrap {
	wire.Build(config.SetupFiber, config.SetupLogging, initialization.ProvideBootstrap)
	return initialization.Bootstrap{}
}
