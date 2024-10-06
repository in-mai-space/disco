package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/in-mai-space/disco/database"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello, world!")
		return nil
	})

	app.Listen(":3000")
}
