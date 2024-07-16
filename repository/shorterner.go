package repository

import (
	"context"
	. "peterchu999/url-shorterner/database"
	. "peterchu999/url-shorterner/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "url"

func GetUrlData(shortUrl string) (URL, error) {
	collection := MongoDB.Collection(COLLECTION_NAME)
	res := collection.FindOne(context.TODO(), bson.M{"short_url": shortUrl})
	var url URL
	res.Decode(&url)
	return url, res.Err()
}

func CreateUrlData(url URL) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "short_url", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	collection := MongoDB.Collection(COLLECTION_NAME)
	collection.Indexes().CreateOne(context.TODO(), indexModel)
	_, err := collection.InsertOne(context.TODO(), url)
	if err != nil {
		return err
	}
	opsUpsert := options.Update().SetUpsert(true)
	updateBson := bson.D{
		{Key: "$inc", Value: bson.D{{Key: "count", Value: 1}}}, {Key: "$setOnInsert", Value: bson.D{{Key: "name", Value: "url"}}}}
	_, err2 := MongoDB.Collection(COLLECTION_NAME+"_indexes").UpdateOne(context.TODO(), bson.M{"name": "url"}, updateBson, opsUpsert)

	return err2
}

func GetCurrentCount() (int, error) {
	collection := MongoDB.Collection(COLLECTION_NAME + "_indexes")
	res := collection.FindOne(context.TODO(), bson.M{"name": "url"})

	if res.Err() == mongo.ErrNoDocuments {
		return 0, nil
	} else if res.Err() != nil {
		return 0, res.Err()
	}
	var index URLIndex
	err := res.Decode(&index)
	return index.Count, err
}
