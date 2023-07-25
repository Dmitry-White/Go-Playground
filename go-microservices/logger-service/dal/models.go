package dal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (le *LogEntry) GetAll() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	collection := db.Database("logs").Collection("logs")

	filter := bson.D{}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	logs := []*LogEntry{}
	for cursor.Next(ctx) {
		item := LogEntry{}

		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		logs = append(logs, &item)
	}

	return logs, err
}

func (le *LogEntry) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	collection := db.Database("logs").Collection("logs")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": docId}
	result := collection.FindOne(ctx, filter)

	item := LogEntry{}
	decodeErr := result.Decode(&item)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &item, nil
}

func (le *LogEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	collection := db.Database("logs").Collection("logs")

	docId, err := primitive.ObjectIDFromHex(le.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": docId}
	updatedEntry := bson.D{
		{Key: "name", Value: le.Name},
		{Key: "data", Value: le.Data},
		{Key: "updated_at", Value: time.Now()},
	}
	payload := bson.D{
		{Key: "$set", Value: &updatedEntry},
	}
	result, err := collection.UpdateOne(ctx, filter, payload)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (le *LogEntry) Insert(entry LogEntry) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	collection := db.Database("logs").Collection("logs")

	payload := LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection.InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (le *LogEntry) DropAll() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	collection := db.Database("logs").Collection("logs")

	err := collection.Drop(ctx)
	if err != nil {
		return "", err
	}

	return "logs", nil
}
