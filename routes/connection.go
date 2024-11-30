package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {
	MONGODB_CONNECT_URL := "mongodb://foo:bar@localhost:27017"
	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, connectErr := mongo.Connect(options.Client().ApplyURI(MONGODB_CONNECT_URL))
	if connectErr != nil {
		log.Printf("\nConnect Error: %v\n", connectErr.Error())

	}
	fmt.Println("Connected to MongoDB client...")
	return client
}

var Client *mongo.Client = DBinstance()

func openCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("calorieDB").Collection(collectionName)
	return collection
}
