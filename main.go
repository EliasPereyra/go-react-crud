package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Here is the data requested",
		})
	})

	app.Listen(":" + port)
	fmt.Println("Server on port 3000")
}
