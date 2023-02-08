package main

import (
	"github.com/EliasPereyra/go-react-crud/src/config"
	"github.com/EliasPereyra/go-react-crud/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	config.ConnectMongoDB()
	
	app.Use(cors.New())

	routes.UserRoute(app)

	app.Listen(":")
}
