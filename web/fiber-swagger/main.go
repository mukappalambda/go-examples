package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/mukappalambda/go-examples/web/fiber-swagger/docs"
)

// @title			Fiber Example API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	app := fiber.New()
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/users/", GetAllUsers)
	log.Fatal(app.Listen(":8080"))
}

// GetAllUsers is a function to get all users data from database
//
//	@Summary		Get all users
//	@Description	Get all users
//	@Tags			users
//	@Router			/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	return c.Send([]byte("All books"))
}
