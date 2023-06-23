package initialization

import (
	"cleansoftware.io/ddd/fiber/seed/internal/config"
	"cleansoftware.io/ddd/fiber/seed/internal/products/domain/ports"
	"github.com/gofiber/fiber/v2"
)

type Bootstrap struct {
	Logger ports.Logger
	App    *fiber.App
}

func ProvideBootstrap(app *fiber.App, logger ports.Logger) Bootstrap {
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
