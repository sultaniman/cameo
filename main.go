package main

import (
	"github.com/gofiber/fiber"
	"github.com/imanhodjaev/cameo/cameo"
)

func main() {
	app := fiber.New()
	cameo.LoadConfig()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})
	app.Use(func() {})

	_ = app.Listen(3000)
}
