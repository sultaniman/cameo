package cameo

import (
	"github.com/gofiber/csrf"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
	"github.com/imanhodjaev/cameo/cameo/config"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"net/url"
)

type App struct {
	Path   string
	Config *config.Config
}

func Serve(app *App) error {
	server := fiber.New(&fiber.Settings{
		Views:                 createTemplateEngine(),
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

func createTemplateEngine() *html.Engine {
	return &html.Engine{
		Templates: template.Must(template.New("form").Parse(FormTemplate)),
	}
}
