package controllers

import (
	"math"
	"time"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/mike-neutron/go_currency_app/src/initializers"
	"github.com/mike-neutron/go_currency_app/src/models"
)

func GetRatesByDate(c *fiber.Ctx) error {
	// Define result array
	var result []models.Rate

	// Make rates query
	dateString := c.Params("date");
	initializers.DB.Where("date = ?", dateString).Find(&result)	
		// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{"rates": result},
	})
}

func ExchangeRate(c *fiber.Ctx) error {
	from  := c.Params("from")
	to    := c.Params("to")
	value := c.Params("value")
	valueFloat64, _ := strconv.ParseFloat(value, 32)
	valueFloat32 := float32(valueFloat64)

	currentTime := time.Now()
	currentDate := currentTime.Format("01-02-2006")

	var fromRow models.Rate
	var toRow models.Rate

	if (from != "RUB") {
		errFrom := initializers.DB.Where("date = ? AND code = ?", currentDate, from).First(&fromRow)
	
		if (errFrom != nil) {
			if errFrom.RowsAffected != 1 {
				// Return response
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status": "fail",
					"data": fiber.Map{
						"error": "Wrong FROM currency",
					},
				})
			}
		}
	} else {
		fromRow = models.Rate{
			Code: "RUB",
			Rate: 1,
		}
	}
	

	if (to != "RUB") {
		errTo := initializers.DB.Where("date = ? AND code = ?", currentDate, to).First(&toRow)
	
		if (errTo != nil) {
			if errTo.RowsAffected != 1 {
				// Return response
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"status": "fail",
					"data": fiber.Map{
						"error": "Wrong TO currency",
					},
				})
			}
		}
	} else {
		toRow = models.Rate{
			Code: "RUB",
			Rate: 1,
		}
	}


	// Calculate
	var result float32 = valueFloat32 * fromRow.Rate / toRow.Rate
	resultForReturn := math.Floor(math.Abs(float64(result)) * 100) / 100

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"result": resultForReturn,
		},
	})
}