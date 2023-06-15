package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/api/stocks/current", func(c *fiber.Ctx) error {
		return c.SendString("GET /api/stocks/current method")
	})

	app.Get("/api/stocks/date/:date", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("date"))
	})

	app.Get("/api/stocks/exchange/:from/:to", func(c *fiber.Ctx) error {
		return c.SendString("from: "+c.Params("from")+ ", to: "+c.Params("to"))
	})

    app.Listen(":8080")
}
