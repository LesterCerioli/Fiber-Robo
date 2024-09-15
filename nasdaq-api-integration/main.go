package main

import (
	"fmt"
	"log"
	"nasdaq-api-integration/config"
	"nasdaq-api-integration/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load API config (Nasdaq API key)
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	
	app := fiber.New()

	
	app.Get("/stock/:symbol", func(c *fiber.Ctx) error {
		symbol := c.Params("symbol")
		stockData, err := services.FetchStockData(symbol)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(stockData)
	})

	
	app.Get("/exchange/:currency", func(c *fiber.Ctx) error {
		currency := c.Params("currency")
		rate, err := services.GetExchangeRate(currency)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(fiber.Map{"currency": currency, "rate": rate})
	})

	// Start the app
	log.Fatal(app.Listen(":3000"))
}
	