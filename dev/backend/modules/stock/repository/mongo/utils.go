package stock_repository

import (
	"context"
	"errors"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Seems really of to define this here
const ItemsPerPage int64 = 12

func (repo *Repository) HasProduct(ctx context.Context, id string) bool {
	err := repo.collection.FindOne(ctx, id).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}

func (repo *Repository) TotalItems(ctx context.Context, filter bson.M) (int64, error) {
	total, err := repo.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *Repository) TotalPages(ctx context.Context, filter bson.M) (int64, error) {
	total, err := repo.TotalItems(ctx, filter)
	if err != nil {
		return 0, err
	}

	// Calculate the quantity of pages
	return int64(math.Ceil(float64(total) / float64(ItemsPerPage))), nil
}
