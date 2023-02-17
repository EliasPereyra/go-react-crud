package controllers

import (
	"context"
	"fmt"

	"github.com/EliasPereyra/go-react-crud/src/config"
	"github.com/EliasPereyra/go-react-crud/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollections(config.DB, "users")

func GetUsers(c *fiber.Ctx) error{
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	
	var results []models.User
	if err = cursor.All(c.Context(), &results); err != nil {
		 panic(err)
	}
	fmt.Print(results)
	return c.JSON(results)
}

func GetUser(c *fiber.Ctx) error {
	userId := new(models.User)

	c.ParamsParser(userId)

	result := userCollection.FindOne(context.TODO(), bson.D{{"_id", userId}})

	return c.JSON(result)
}

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
