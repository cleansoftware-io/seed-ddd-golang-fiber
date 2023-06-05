package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func SetupFiber(requestLatency *RequestLatency) *fiber.App {

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		requestLatency.Observe(c.Method(), c.Path(), time.Since(start).Seconds())
		logrus.Infof("%s %s", c.Method(), c.Path())
		logrus.Infof("Status code: %d", c.Response().StatusCode())
		return c.Next()

	})

	return app
}
