package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Panic(err)
	}
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	coll := client.Database("golangAPI").Collection(collectionName)
	return coll
}
