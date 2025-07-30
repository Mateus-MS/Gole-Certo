package stock_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TEMP
func (repo *Repository) ReadRaw(ctx context.Context, filter bson.M) (prod Product, err error) {
	return repo.read(ctx, filter)
}

func (repo *Repository) ReadByID(ctx context.Context, id string) (prod Product, err error) {
	// Convert the received ID into how mongoDB expectes it to be
	var objID primitive.ObjectID
	if objID, err = primitive.ObjectIDFromHex(id); err != nil {
		return prod, err
	}

	// Perform the query
	return repo.read(ctx, bson.M{"_id": objID})
}

func (repo *Repository) ReadByName(ctx context.Context, name string) (prod Product, err error) {
	return repo.read(ctx, bson.M{"name": name})
}

func (repo *Repository) ReadManyFilteredAfterID(ctx context.Context, filter bson.M, lastID string, limit int64) (prods []Product, err error) {

	// 1 - Define the time limit for this operation
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	findOptions := options.Find().SetLimit(limit).SetSort(bson.D{{Key: "_id", Value: 1}})

	// 2 - If received a non empty lastID, add it to the query
	if lastID != "" {

		var objID primitive.ObjectID
		// Convert the received ID into how mongoDB expectes it to be
		if objID, err = primitive.ObjectIDFromHex(lastID); err != nil {
			return prods, err
		}

		// Add the id in the filer, allowing only IDs greater than the one received
		filter["_id"] = bson.M{"$gt": objID}
	}

	// 3 - Perform the query
	var result *mongo.Cursor
	if result, err = repo.collection.Find(ctx, filter, findOptions); err != nil {
		return prods, err
	}
	defer result.Close(ctx)

	// 4 - Decode the result
	if err := result.All(ctx, &prods); err != nil {
		return nil, err
	}
	return prods, nil
}
