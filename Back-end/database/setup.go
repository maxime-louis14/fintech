// database/database.go
package database

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(EnvMongoURI())
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// Getting database collections
func GetSondageCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("qestions-fintech").Collection("Sondage")
	return collection
}

func GetQuestionnaireCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("qestions-fintech").Collection("Questionnaire")
	return collection
}

func GetAdminFormCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("qestions-fintech").Collection("adminform")
	return collection
}
