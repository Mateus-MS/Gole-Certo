package stock_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) DeleteByID(ctx context.Context, id string) error {
	return repo.delete(ctx, bson.M{"_id": id})
}

func (repo *Repository) DeleteByName(ctx context.Context, name string) error {
	return repo.delete(ctx, bson.M{"name": name})
}
