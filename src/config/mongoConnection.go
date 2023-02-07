package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	cfg := getConfig()

	if cfg.MONGODB_URI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MONGODB_URI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var DB *mongo.Client = ConnectMongoDB()

func GetCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-mongo").Collection("users")

	return collection
}
