package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}
	
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	type User struct {
		Name	string
		Lastname	string
		Age		int 
	} 

	new_user := User{Name: "mark", Lastname: "schmidt", Age: 25}

	coll := client.Database("go-mongo").Collection("users")
	result, err := coll.InsertOne(context.TODO(), new_user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	
	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Here is the data requested",
		})
	})

	app.Listen(":" + port)
}
