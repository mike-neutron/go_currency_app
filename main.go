package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mike-neutron/go_currency_app/src/initializers"
	"github.com/mike-neutron/go_currency_app/src/controllers"
	"github.com/mike-neutron/go_currency_app/src/seeders"
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/mike-neutron/go_currency_app/docs"
)

func init() {
	config, err := initializers.LoadConfig(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

// @title Rates API
// @version 1.0
// @description api for getting rates, exchange rates
// @BasePath /
func main() {
    app := fiber.New()
	app.Use(logger.New())

	seeders.SeedRates()
    app.Get("/", controllers.HelloWorld)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/api/date/:date<datetime(2006\\-01\\-02)>", controllers.GetRatesByDate)
	app.Get("/api/exchange/:from/:to/:value<float>", controllers.ExchangeRate)
	log.Fatal(app.Listen(":8080"))
}
