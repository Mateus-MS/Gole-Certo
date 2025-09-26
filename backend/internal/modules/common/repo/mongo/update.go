package generic_repository_mongo

import (
	"context"

	generic_repository "alves.com/modules/common/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *GenericRepository[T]) Update(ctx context.Context, filter bson.M, updateSet bson.M) error {
	result, err := repo.Collection.UpdateOne(ctx, filter, updateSet)

	if result.MatchedCount == 0 {
		return generic_repository.ErrItemInexistent
	}

	if err != nil {
		return err
	}

	return nil
}
