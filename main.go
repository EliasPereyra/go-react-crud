package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Here is the data requested",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server on port 3000")
}
