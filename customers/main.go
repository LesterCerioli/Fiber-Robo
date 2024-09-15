package main

import(
	"github.com/gofiber/fiber/v2"
    "customers/routes"
    "log"
)

func nain() {
    app := fiber.new()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}