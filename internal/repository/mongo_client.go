package repository

import (
	"context"
	_ "context"
	"fmt"
	_ "fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	_ "log"
	_ "time"

	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func InitMongoDB(uri, dbName, collectionName string) {
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("crud_db").Collection("users")
}

func GetMongoCollection() *mongo.Collection {
	return collection
}

func DisconnectMongoDB() {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB")
}
