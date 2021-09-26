package gRPC_Mongo_HTTP_TestTask

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client
var Collections *mongo.Collection

// ConnectDatabase is used to connect the MongoDB database
func ConnectDatabase() {
	log.Println("Database connecting...")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	Client = client
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected.")

	Collections = client.Database("mydb").Collection("book")
}

func DBClose() error {
	ctx := context.Background()
	return Client.Disconnect(ctx)
}
