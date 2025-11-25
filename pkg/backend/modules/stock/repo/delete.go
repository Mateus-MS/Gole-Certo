package stock_repository

import (
	"context"

	stock_error "alves.com/backend/modules/stock/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) Delete(ctx context.Context, filter bson.M) error {
	var result *mongo.DeleteResult

	result, err := repo.Collection.DeleteOne(ctx, filter)

	if result.DeletedCount == 0 {
		return stock_error.ErrStockInexistent
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return repo.Delete(ctx, bson.M{"_id": id})
}
