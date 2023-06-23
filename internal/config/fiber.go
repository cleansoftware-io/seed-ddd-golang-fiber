package config

import (
	"cleansoftware.io/ddd/fiber/seed/internal/domain/ports"
	"github.com/gofiber/fiber/v2"
)

func SetupFiber(logger ports.Logger) *fiber.App {

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		logger.Info("%s %s", c.Method(), c.Path())
		logger.Info("Status code: %d", c.Response().StatusCode())
		return c.Next()

	})

	return app
}
