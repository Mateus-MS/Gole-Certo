package stock_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) UpdateByID(ctx context.Context, prod Product) (err error) {
	// Build the query
	filter := bson.M{"_id": prod.ProductID}

	// Perform the query
	if err = repo.update(ctx, prod, filter); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateByName(ctx context.Context, prod Product) (err error) {
	// Build the query
	filter := bson.M{"name": prod.Name}

	// Perform the query
	if err = repo.update(ctx, prod, filter); err != nil {
		return err
	}

	return nil
}
