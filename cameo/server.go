package cameo

import (
	"github.com/gofiber/csrf"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/imanhodjaev/cameo/cameo/app"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type App struct {
	Path   string
	Config *app.Config
}

func Serve(app *App) error {
	server := fiber.New(&fiber.Settings{
		Views:                 CreateTemplateEngine(),
		DisableStartupMessage: true,
	})

	server.Post("/", app.SendMessage)
	server.Get("/", app.ShowForm)

	server.Use(CheckDomains(app.Config.Domains))
	server.Use(csrf.New())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	return server.Listen(app.Config.Port)
}

func CheckDomains(domains []string) fiber.Handler {
	return func(c *fiber.Ctx) {
		uri, err := url.ParseRequestURI(c.BaseURL())
		if err != nil {
			logrus.Warn("Unable to parse request URI", err)
			c.Next()
			return
		}

		for _, domain := range domains {
			if uri.Host == domain {
				c.Next()
				return
			}
		}

		c.SendStatus(http.StatusBadRequest)
	}
}
