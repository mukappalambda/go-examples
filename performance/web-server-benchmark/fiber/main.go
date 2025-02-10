package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var port = flag.Int("port", 8080, "server port")

func main() {
	flag.Parse()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello from the gofiber server."))
	})
	fmt.Printf("server listening at %d\n", *port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
