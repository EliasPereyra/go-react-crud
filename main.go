package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
		Name			string		`json:"name" form:"name"`
		Lastname	string		`json:"lastName" form:"lastName"`
		Age				int 			`json:"age" form:"age"`
	} 

	coll := client.Database("go-mongo").Collection("users")
	
	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		cursor, err := coll.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}

		var results []User
		if err = cursor.All(c.Context(), &results); err != nil {
			panic(err)
		}

		return c.JSON(results)
	})

	app.Post("/users/:id", func(c *fiber.Ctx) error {
		newUser := new(User)
		c.BodyParser(newUser)

		result, err := coll.InsertOne(context.TODO(), newUser)
		if err != nil {
			panic(err)
		}

		fmt.Println("user created")
		return c.JSON(&fiber.Map{
			"user": result,
		})
	})

	app.Listen(":" + port)
}
