package repositories

import (
	"ciri2-pc-microservice/configs"
	"ciri2-pc-microservice/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ComponentRepository struct{}

var componentCollection *mongo.Collection = configs.GetCollection(configs.DB, "components")

func init() {

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := componentCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}
}

func (c ComponentRepository) BatchInsert(components []interface{}) (*mongo.BulkWriteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var operations []mongo.WriteModel
	for _, component := range components {
		filter := bson.M{"name": component.(models.Component).Name}
		update := bson.M{"$set": component}
		operation := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		operations = append(operations, operation)
	}

	result, err := componentCollection.BulkWrite(ctx, operations)
	return result, err
}

func (c ComponentRepository) FindOne(id string) (interface{}, error) {
	var component models.Component
	objectId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := componentCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&component)
	return component, err
}

func (c ComponentRepository) FindPaginated(pagination models.Pagination) ([]models.Component, int, error) {
	var components []models.Component
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalCount, err := componentCollection.CountDocuments(ctx, pagination.Filter)
	if err != nil {
		return nil, 0, err
	}

	if totalCount == 0 {
		return components, 0, nil
	}

	if pagination.Skip > int(totalCount) {
		pagination.Skip = int(totalCount)
	}

	options := options.Find()
	options.SetSkip(int64((pagination.Skip - 1) * pagination.Limit))
	options.SetLimit(int64(pagination.Limit))
	options.SetSort(pagination.Sort)

	cursor, err := componentCollection.Find(ctx, pagination.Filter, options)
	if err != nil {
		return nil, 0, err
	}
	if err = cursor.All(context.TODO(), &components); err != nil {
		panic(err)
	}

	return components, int(totalCount), nil
}

func (c ComponentRepository) CheckAlreadyExistingComponents(componentNames []string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var components []models.Component

	cursor, err := componentCollection.Find(ctx, bson.M{"name": bson.M{"$in": componentNames}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &components)
	if err != nil {
		return nil, err
	}

	existingComponents := make([]string, len(components))
	for i, component := range components {
		existingComponents[i] = component.Name
	}

	return existingComponents, nil
}
