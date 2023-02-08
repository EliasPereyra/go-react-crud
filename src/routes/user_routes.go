package routes

import (
	"github.com/EliasPereyra/go-react-crud/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		message := "Welcome to the server. Later you'll be redirected to the UI."
		
		return c.JSON(message)
		})
	
	app.Get("/users", controllers.GetUsers)
}
