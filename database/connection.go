package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func catchError(err error, msg string) {
	if err != nil {
		if msg == "" {
			log.Fatal(err)
		} else {
			log.Fatal(msg)
		}
	}
}

func InitMongoDB() *mongo.Client {
	err := godotenv.Load(".env")
	catchError(err, "Error loading .env file")
	MongodbURL := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongodbURL))
	catchError(err, "")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	catchError(err, "")
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = InitMongoDB()

func OpenCollection(client *mongo.Client, collectName string) *mongo.Collection {

	var collection *mongo.Collection = (*mongo.Collection)(client.Database(os.Getenv("DATABASE_NAME")).Collection(collectName))
	return collection
}
