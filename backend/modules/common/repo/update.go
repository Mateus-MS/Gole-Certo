package generic_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *GenericRepository[T]) Update(ctx context.Context, filter bson.M, updateSet bson.M) error {
	result, err := repo.collection.UpdateOne(ctx, filter, updateSet)

	if result.MatchedCount == 0 {
		return ErrItemInexistent
	}

	if err != nil {
		return err
	}

	return nil
}
