package main

import (
	"log"
	"payment-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	dsn := "host=localhost user=postgres password=your_password dbname=payment_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB.AutoMigrate(&models.Customer{}, &models.Payment{})
}

func main() {
	app := fiber.New()

	// Initialize the database
	initDatabase()

	// Define routes
	app.Post("/payment", createPayment)
	app.Get("/payment/:id", getPayment)

	log.Fatal(app.Listen(":3001"))
}

// Handlers for the API
func createPayment(c *fiber.Ctx) error {
	payment := new(models.Payment)
	if err := c.BodyParser(payment); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	DB.Create(&payment)
	return c.JSON(payment)
}

func getPayment(c *fiber.Ctx) error {
	id := c.Params("id")
	var payment models.Payment
	if result := DB.First(&payment, id); result.Error != nil {
		return c.Status(404).SendString(result.Error.Error())
	}
	return c.JSON(payment)
}
