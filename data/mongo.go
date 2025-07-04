package data

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// comment "TaskCollection" is a global variable that holds the MongoDB collection for tasks.
var TaskCollection *mongo.Collection

func ConnectMongoDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("MONGODB_URI environment variable is not set")
	}

	// comment Create a new client and connect to the server
	clientOptions := options.Client().ApplyURI(uri)
	// comment WithTimeout sets a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "taskdb" // default value
	}

	collectionName := os.Getenv("COLLECTION_NAME")
	if collectionName == "" {
		collectionName = "tasks" // default value
	}

	TaskCollection = client.Database(dbName).Collection(collectionName)
}
