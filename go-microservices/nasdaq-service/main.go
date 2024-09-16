package main

import (
	"log"
	"nasdaq-service/models"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	dsn := "host=localhost user=postgres password=your_password dbname=nasdaq_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB.AutoMigrate(&models.Stock{})
}

func main() {
	app := fiber.New()

	// Initialize the database
	initDatabase()

	// Define routes
	app.Get("/stock/:symbol", getStockData)

	log.Fatal(app.Listen(":3002"))
}

// Nasdaq API integration
func getStockData(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	client := resty.New()

	// Simulate Nasdaq API request
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("https://api.nasdaq.com/api/quote/" + symbol + "/info")

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Store in the database
	stock := models.Stock{
		Symbol: symbol,
		Price:  extractPriceFromNasdaq(resp.Body()), // Assume a function to parse the price
	}
	DB.Create(&stock)

	return c.JSON(stock)
}

func extractPriceFromNasdaq(body []byte) float64 {
	// Example parser logic here to extract stock price from response body
	return 100.0 // Placeholder
}
