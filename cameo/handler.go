package cameo

import (
	"github.com/gofiber/fiber"
)

func (a *App) SendMessage(c *fiber.Ctx) {
	var data *fiber.Map
	config := a.Config
	message := Message{
		Subject: c.FormValue("subject"),
		Body:    c.FormValue("body"),
	}

	if message.IsValid() {
		config.Mailer.SendAsync(&message, &config.GPG)
		data = a.getData(&Message{}, "Message sent.", false)
	} else {
		data = a.getData(&message, "Please fill out subject and message", true)
	}

	_ = c.Render("form", data)
}

func (a *App) ShowForm(c *fiber.Ctx) {
	_ = c.Render("form", a.getData(&Message{}, "", false))
}

func (a *App) getData(message *Message, note string, isError bool) *fiber.Map {
	return &fiber.Map{
		"Message": message,
		"Title":   a.Config.FormTitle,
		"Note":    note,
		"IsError": isError,
		"Version": a.Config.Version,
		"Favicon": Favicon,
	}
}
