package stock_repository

import (
	"context"

	stock_error "alves.com/backend/modules/stock/errors"
	stock_model "alves.com/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) Update(ctx context.Context, filter bson.M, updateSet bson.M) error {
	result, err := repo.Collection.UpdateOne(ctx, filter, updateSet)

	if result.MatchedCount == 0 {
		return stock_error.ErrStockInexistent
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateByID(ctx context.Context, stock stock_model.StockEntity) error {
	update := bson.M{
		"$set": bson.M{
			"name":     stock.Name,
			"quantity": stock.Quantity,
		},
	}

	result, err := repo.Collection.UpdateOne(ctx, bson.M{"_id": stock.ID}, update)

	if result.MatchedCount == 0 {
		return stock_error.ErrStockInexistent
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AtomicDecreaseStockByID(ctx context.Context, prodID primitive.ObjectID, quantity int) (*mongo.UpdateResult, error) {
	return repo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": prodID, "quantity": bson.M{"$gte": quantity}},
		bson.M{"$inc": bson.M{"quantity": -quantity}},
	)
}
