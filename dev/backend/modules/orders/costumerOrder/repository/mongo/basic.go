package costumerOrder_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) read(ctx context.Context, filter bson.M) (ord Order, err error) {
	if err = repo.collection.FindOne(ctx, filter).Decode(&ord); err != nil {
		return ord, ErrOrderNotFound
	}

	return ord, nil
}
