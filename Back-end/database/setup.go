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
	defer cancel() // Appel de la fonction de cancellation à la fin de la portée

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping la base de données
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, recette string) *mongo.Collection {
	collection := client.Database("qestions-fintech").Collection("Sondage")
	return collection
}
