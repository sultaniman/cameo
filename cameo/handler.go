package cameo

import (
	"github.com/gofiber/fiber"
	"github.com/imanhodjaev/cameo/cameo/app"
)

func (a *App) SendMessage(c *fiber.Ctx) {
	var data *fiber.Map
	message := app.Message{
		Subject: c.FormValue("subject"),
		Body:    c.FormValue("body"),
	}

	if message.IsValid() {
		app.SendAsync(&message, a.Config)
		data = a.getData(&app.Message{}, "Message sent.", false)
	} else {
		data = a.getData(&message, "Please fill out subject and message of at least 2 words", true)
	}

	_ = c.Render("form", data)
}

func (a *App) ShowForm(c *fiber.Ctx) {
	_ = c.Render("form", a.getData(&app.Message{}, "", false))
}

func (a *App) getData(message *app.Message, note string, isError bool) *fiber.Map {
	return &fiber.Map{
		"Message": message,
		"Title":   a.Config.FormTitle,
		"Note":    note,
		"IsError": isError,
		"Version": a.Config.Version,
		"Favicon": Favicon,
	}
}
