package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})
	app.Use()

	_ = app.Listen(3000)
}
