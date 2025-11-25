package stock_repository

import (
	"context"

	stock_error "alves.com/backend/modules/stock/errors"
	stock_model "alves.com/backend/modules/stock/model"
)

func (repo *Repository) Create(ctx context.Context, stock stock_model.StockEntity) error {
	_, err := repo.Collection.InsertOne(ctx, stock)

	if err != nil {
		return stock_error.ErrStockAlreadyExists
	}

	return nil
}
