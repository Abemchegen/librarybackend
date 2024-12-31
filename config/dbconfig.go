package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(url string, DBname string) (mongo.Database, mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
		return mongo.Database{}, mongo.Client{}, err
	}

	fmt.Println("Connected to MongoDB!")

	database := client.Database(DBname)

	return *database, *client, nil
}

func CreateCollection(database mongo.Database, collectionName string) mongo.Collection {
	collection := database.Collection(collectionName)
	return *collection
}

func CloseDB(client mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
