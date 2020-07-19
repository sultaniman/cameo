package cameo

import (
	"github.com/gofiber/fiber"
)

type AppConfig struct {
	Path   string
	Config *Config
}

func serve(config *AppConfig) error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Use(func() {})

	return app.Listen(3000)
}
