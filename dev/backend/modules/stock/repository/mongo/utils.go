package stock_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) HasProduct(ctx context.Context, id string) bool {
	err := repo.collection.FindOne(ctx, id).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}
