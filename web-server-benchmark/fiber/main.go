package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello from the gofiber server."))
	})
	log.Fatal(app.Listen(":8082"))
}
