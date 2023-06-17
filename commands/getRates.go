package main

import (
	"log"
	"github.com/mike-neutron/go_currency_app/src/initializers"
	"github.com/mike-neutron/go_currency_app/src/seeders"
)

func init() {
	config, err := initializers.LoadConfig("/app/.env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	seeders.SeedRates()
}