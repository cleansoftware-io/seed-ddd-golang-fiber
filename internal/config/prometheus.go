package config

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupPrometheus(app *fiber.App) {
	fiberPrometheus := fiberprometheus.New("seed_fiber_golang_server")
	fiberPrometheus.RegisterAt(app, "/metrics")
	app.Use(fiberPrometheus.Middleware) //use default metrics from fiberprometheus
}
