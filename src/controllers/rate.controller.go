package controllers

import (
	"math"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/mike-neutron/go_currency_app/src/initializers"
	"github.com/mike-neutron/go_currency_app/src/models"
)

//	@Summary		Get rates by date
//	@Description	Get rates by date
//	@Produce		json
//	@Param			date	path		string	true "Date of rate"	"Date" default(2023-06-17)
//	@Success		200	{object}	GetRatesByDateResponse
//	@Router			/api/date/{date} [get]
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

//	@Summary		ExchangeRate
//	@Description	ExchangeRate
//	@Produce		json
//	@Param			from	path		string	true	"From currency" default(USD)
//	@Param			to	path		string	true	"To currency" default(RUB)
//	@Param			value	path		number	true	"Value" default(100)
//	@Success		200	{object}	ExchangeRateResponse
//	@Router			/api/exchange/{from}/{to}/{value} [get]
func ExchangeRate(c *fiber.Ctx) error {
	from  := c.Params("from")
	to    := c.Params("to")
	value := c.Params("value")
	valueFloat64, _ := strconv.ParseFloat(value, 32)

	var fromRow, toRow models.Rate
	rubRate := models.Rate{
		Code: "RUB",
		Rate: 1,
	}

	if (from != "RUB") {
		errFrom := initializers.DB.Where("code = ?", from).Order("date desc").First(&fromRow)
	
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
		fromRow = rubRate
	}
	

	if (to != "RUB") {
		errTo := initializers.DB.Where("code = ?", to).Order("date desc").First(&toRow)
	
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
		toRow = rubRate
	}


	// Calculate
	var result float64 = valueFloat64 * float64(fromRow.Rate) / float64(toRow.Rate)
	resultForReturn := math.Floor(math.Abs(float64(result)) * 10000) / 10000

	// Return response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"result": resultForReturn,
		},
	})
}

type GetRatesByDateResponse struct {
	Status string `json:"status" example:"success"`
	Data struct{
		Rates []models.RateSwagger `json:"rates"`
	} `json:"data"`
}

type ExchangeRateResponse struct {
	Status string `json:"status" example:"success"`
	Data struct{
		Result float32 `json:"result" example:"10.91"`
	} `json:"data"`
}
