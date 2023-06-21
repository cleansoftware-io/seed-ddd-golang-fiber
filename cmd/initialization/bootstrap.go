package initialization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.con/tgarcia/seed-ddd-golang-fiber/internal/config"
)

type Bootstrap struct {
	Logger   *logrus.Logger
	Registry *prometheus.Registry
	App      *fiber.App
}

func NewBootstrap() (*Bootstrap, error) {

	logger := config.SetupLogging()
	app := config.SetupFiber(logger)
	config.SetupPrometheus(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		msg := "Green"
		return c.SendString(msg)
	})

	return &Bootstrap{
		App:    app,
		Logger: logger,
	}, nil
}

// Start inicia el servidor HTTP.
func (b *Bootstrap) Start() error {
	return b.App.Listen(":3000")
}
