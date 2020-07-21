package cameo

import (
	"github.com/gofiber/csrf"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
)

type App struct {
	Path   string
	Config *Config
}

type Message struct {
	Subject string
	Body 	string
}

func Serve(app *App) error {
	server := fiber.New(&fiber.Settings{
		Views: html.New("./templates", ".html"),
	})

	server.Post("/", app.SendMessage)
	server.Get("/", app.ShowForm)

	server.Use(csrf.New())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	return server.Listen(app.Config.Port)
}
