//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.con/tgarcia/seed-ddd-golang-fiber/cmd/initialization"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/config"
)

func InitializeApplication() initialization.Bootstrap {
	wire.Build(config.SetupFiber, config.SetupLogging, initialization.ProvideBootstrap)
	return initialization.Bootstrap{}
}
