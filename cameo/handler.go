package cameo

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber"
)

func (a *App) SendMessage(c *fiber.Ctx) {
	fmt.Println("POST…")
	data := fiber.Map{
		"Message": Message{
			Subject: c.FormValue("Subject"),
			Body: c.FormValue("Body"),
		},
		"Title": a.Config.FormTitle,
	}

	_ = c.Render("form", data)
}


func (a *App) ShowForm(c *fiber.Ctx) {
	spew.Dump("GET…")
	_ = c.Render("form", fiber.Map{
		"Title": a.Config.FormTitle,
		"Message": Message{},
	})
}
