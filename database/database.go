package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

var RedisClient *redis.Client

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

func ConnectRedis() error {
	redisUri := os.Getenv("REDIS_URI")
	redisPassword := os.Getenv("REDIS_PASSSWORD")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	return nil
}
