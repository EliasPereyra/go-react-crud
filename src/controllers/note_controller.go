package controllers

import (
	"context"

	"github.com/EliasPereyra/go-react-crud/src/config"
	"github.com/EliasPereyra/go-react-crud/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection = config.GetCollections(config.DB, "notes")

func GetNotes(c *fiber.Ctx) error {
	cursor, err := noteCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var allNotes []models.Note
	if err = cursor.All(c.Context(), *&allNotes); err != nil {
		panic(err)
	}
	println(allNotes)

	return c.JSON(allNotes)
}
