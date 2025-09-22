package generic_repository

import (
	"context"
	"errors"

	generic_persistent "alves.com/modules/common/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *GenericRepository[T]) Read(ctx context.Context, filter bson.M) (generic_persistent.IPersistent, error) {
	var item T

	err := repo.Collection.FindOne(ctx, filter).Decode(&item)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return item, ErrItemInexistent
		}
	}

	return item, nil
}
