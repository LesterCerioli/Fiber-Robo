package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"nasdaq-integration/models"
)

var DB *gorm.DB

func initDatabase() {
	dsn := "host=localhost user=postgres password=your_password dbname=your_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB.AutoMigrate(&models.Customer{}, &models.Payment{})
	fmt.Println("Database connected successfully!")
}

func main() {
	app := fiber.New()

	initDatabase()

	registerNasdaqRoutes(app)
	registerPaymentRoutes(app)
	registerCustomerRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
