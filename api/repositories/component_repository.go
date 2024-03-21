package repositories

import (
	"ciri2-pc-microservice/configs"
	"ciri2-pc-microservice/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var componentCollection *mongo.Collection = configs.GetCollection(configs.DB, "components")

func BatchInsert(components []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := componentCollection.InsertMany(ctx, components)
	return result, err
}

func FindOne(id string) (interface{}, error) {
	var component models.Component
	objectId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := componentCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&component)
	return component, err
}

func FindPaginated(pagination models.Pagination) ([]models.Component, error) {
	var components []models.Component
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64((pagination.Skip - 1) * pagination.Limit))
	options.SetLimit(int64(pagination.Limit))
	options.SetSort(pagination.Sort)

	cursor, err := componentCollection.Find(ctx, pagination.Filter, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &components); err != nil {
		panic(err)
	}

	return components, nil
}
