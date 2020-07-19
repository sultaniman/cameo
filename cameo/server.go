package cameo

import (
	"github.com/gofiber/fiber"
)

type AppConfig struct {
	Path   string
	Config *Config
}

func Serve(config *AppConfig) error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})


	return app.Listen(config.Config.Port)
}
