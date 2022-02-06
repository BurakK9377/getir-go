package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration model
type Configuration struct {
	MongoDb                      string
	MongoDbConnectionUri         string
	MongoDbRecordsCollectionName string
}

// ConnectDB : This is service function to connect mongoDB
func ConnectDB() *mongo.Collection {
	config := GetConfiguration()
	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoDbConnectionUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database(config.MongoDb).Collection(config.MongoDbRecordsCollectionName)

	return collection
}

// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {
	_ = godotenv.Load(".env")

	configuration := Configuration{
		os.Getenv("MONGO_DB"),
		os.Getenv("MONGO_DB_CONNECTION_URI"),
		os.Getenv("MONGO_DB_RECORDS_COLLECTION_NAME"),
	}

	return configuration
}
