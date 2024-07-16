package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectMonggoDB() error {

	clientOptions := options.Client()

	uri := os.Getenv("MONGO_DB_URI")
	clientOptions.ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	MongoDB = client.Database("url-shorterner")
	err = client.Ping(context.TODO(), nil)

	return err
}
