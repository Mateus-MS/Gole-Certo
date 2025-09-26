package generic_repository_mongo

import (
	"context"
	"errors"

	generic_persistent "alves.com/modules/common/model"
	generic_repository "alves.com/modules/common/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *GenericRepository[T]) Read(ctx context.Context, filter bson.M) (generic_persistent.IPersistent, error) {
	var item T

	err := repo.Collection.FindOne(ctx, filter).Decode(&item)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return item, generic_repository.ErrItemInexistent
		}
	}

	return item, nil
}
