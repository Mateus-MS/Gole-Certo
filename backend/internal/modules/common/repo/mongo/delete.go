package generic_repository_mongo

import (
	"context"

	generic_repository "alves.com/modules/common/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *GenericRepository[T]) Delete(ctx context.Context, filter bson.M) error {
	var result *mongo.DeleteResult

	result, err := repo.Collection.DeleteOne(ctx, filter)

	if result.DeletedCount == 0 {
		return generic_repository.ErrItemInexistent
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *GenericRepository[T]) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return repo.Delete(ctx, bson.M{"_id": id})
}
