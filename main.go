package main

import (
	"context"
	"fmt"

	"github.com/EliasPereyra/go-react-crud/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	app := fiber.New()
	config.ConnectMongoDB()

	type User struct {
		_id				primitive.ObjectID `bson:"_id"`
		Name			string		`json:"name" form:"name"`
		Lastname	string		`json:"lastName" form:"lastName"`
		Age				int 			`json:"age" form:"age"`
	} 

	var userCollection *mongo.Collection = config.GetCollections(config.DB, "users")
	
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		message := "Welcome to the server. Later you'll be redirected to the UI."
		
		return c.JSON(message)
		})

	app.Get("/users", func(c *fiber.Ctx) error {
	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []User
	 	if err = cursor.All(c.Context(), &results); err != nil {
	 		panic(err)
	 	}
	 	fmt.Print(results)
	 	return c.JSON(results)
	})

	// app.Post("/users/:id", func(c *fiber.Ctx) error {
	// 	newUser := new(User)
	// 	c.BodyParser(newUser)

	// 	result, err := coll.InsertOne(context.TODO(), newUser)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("user created")
	// 	return c.JSON(&fiber.Map{
	// 		"user": result,
	// 	})
	// })

	// app.Put("/users/:id", func(c *fiber.Ctx) error {
	// 	// take the id from the user
	// 	userId := c.Body()
	// 	id := coll.FindOne(context.TODO(), bson.D{{"_id", userId}})
		
	// 	filter := bson.D{{"_id", id}}

	// 	update := bson.D{{"$set", filter}}
		
	// 	return c.JSON(update)
	// })

	// app.Delete("/users/:id", func(c *fiber.Ctx) error {
	// 	// delete a user
	// 	user := new(User)
	// 	c.BodyParser(user)

	// 	filter := bson.D{{"Name", user.Name}}

	// 	result, err := coll.DeleteOne(context.TODO(), filter)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("User deleted", result)
	// 	return c.JSON("")
	// })

	app.Listen(":")
}
