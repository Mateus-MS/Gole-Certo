package stock_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (repo *Repository) ReadManyPaged(ctx context.Context, filter bson.M, page int64, limit int64, ascending bool) (prods []Product, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	skip := (page - 1) * limit
	sortOrder := 1
	if !ascending {
		sortOrder = -1
	}

	findOptions := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetSort(bson.D{{Key: "price", Value: sortOrder}})

	cursor, err := repo.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &prods); err != nil {
		return nil, err
	}

	return prods, nil
}
