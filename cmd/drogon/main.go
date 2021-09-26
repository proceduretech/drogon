package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/proceduretech/drogon/platform"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	platform.InitializeDatabase()

	log.Fatal(app.Listen(":3001"))

	defer platform.GetDatabase().Close()
}
