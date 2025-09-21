package generic_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *GenericRepository[T]) Delete(ctx context.Context, filter bson.M) error {
	var result *mongo.DeleteResult

	result, err := repo.collection.DeleteOne(ctx, filter)

	if result.DeletedCount == 0 {
		return ErrItemInexistent
	}

	if err != nil {
		return err
	}

	return nil
}
