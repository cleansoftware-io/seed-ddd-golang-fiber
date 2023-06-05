package internal

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.con/tgarcia/seed-golang-server/internal/config"
)

type Bootstrap struct {
	Logger   *logrus.Logger
	Registry *prometheus.Registry
	App      *fiber.App
}

func NewBootstrap() (*Bootstrap, error) {

	logger := config.SetupLogging()

	requestLatency := config.NewRequestLatency()

	registry := config.SetupPrometheus(requestLatency)

	app := config.SetupFiber(requestLatency)
	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})
	app.Get("/metrics", func(c *fiber.Ctx) error {
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		return c.SendString("Hello, World 👋!")
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	return &Bootstrap{
		App:      app,
		Logger:   logger,
		Registry: registry,
	}, nil
}

// Start inicia el servidor HTTP.
func (b *Bootstrap) Start() error {
	return b.App.Listen(":3000")
}
