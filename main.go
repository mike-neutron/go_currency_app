package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/api/stocks/current", func(c *fiber.Ctx) error {
		return c.SendString("GET /api/stocks/current method")
	})

	app.Get("/api/stocks/date/date", func(c *fiber.Ctx) error {
		return c.SendString("GET /api/stocks/date/date method")
	})

	app.Get("/api/stocks/exchange/USD/RUB", func(c *fiber.Ctx) error {
		return c.SendString("GET /api/stocks/exchange/USD/RUB method")
	})

    app.Listen(":8080")
}
