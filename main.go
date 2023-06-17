package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mike-neutron/go_currency_app/src/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
    app := fiber.New()
	app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/api/stocks/date/:date", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("date"))
	})

	app.Get("/api/stocks/exchange/:from/:to", func(c *fiber.Ctx) error {
		return c.SendString("from: "+c.Params("from")+ ", to: "+c.Params("to"))
	})

	log.Fatal(app.Listen(":8080"))
}
