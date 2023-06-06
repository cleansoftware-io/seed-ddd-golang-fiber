package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupFiber(logger *logrus.Logger) *fiber.App {

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		logger.Infof("%s %s", c.Method(), c.Path())
		logger.Infof("Status code: %d", c.Response().StatusCode())
		return c.Next()

	})

	return app
}
