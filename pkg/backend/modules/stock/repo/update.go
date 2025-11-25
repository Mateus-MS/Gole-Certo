package stock_repository

import (
	"context"

	stock_error "alves.com/backend/modules/stock/errors"
	"go.mongodb.org/mongo-driver/bson"
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
