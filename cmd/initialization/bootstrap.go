package initialization

import (
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Bootstrap struct {
	Logger *logrus.Logger
	App    *fiber.App
}

func ProvideBootstrap(app *fiber.App, logger *logrus.Logger) Bootstrap {
	config.SetupPrometheus(app)
	setupHealth(app)
	return Bootstrap{
		App:    app,
		Logger: logger,
	}
}

func setupHealth(app *fiber.App) fiber.Router {
	return app.Get("/health", func(c *fiber.Ctx) error {
		msg := "Green"
		return c.SendString(msg)
	})
}

func (b *Bootstrap) Start() error {
	return b.App.Listen(":3000")
}
