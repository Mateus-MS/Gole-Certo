package stock_repository

import (
	"context"
	"errors"

	stock_error "alves.com/backend/modules/stock/errors"
	stock_model "alves.com/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) Read(ctx context.Context, filter bson.M) (stock_model.StockEntity, error) {
	var stock stock_model.StockEntity

	err := repo.Collection.FindOne(ctx, filter).Decode(&stock)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return stock, stock_error.ErrStockInexistent
		}
	}

	return stock, nil
}

func (repo *Repository) ReadByName(ctx context.Context, name string) (*stock_model.StockEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"name": name})

	if err != nil {
		if errors.Is(err, stock_error.ErrStockInexistent) {
			return &stock_model.StockEntity{}, stock_error.ErrStockInexistent
		} else {
			return &stock_model.StockEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	return &userGeneric, nil
}
