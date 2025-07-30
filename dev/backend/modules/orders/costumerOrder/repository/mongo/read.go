package costumerOrder_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByID(ctx context.Context, id string) (ord Order, err error) {
	return repo.read(ctx, bson.M{"_id": id})
}
