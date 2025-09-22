package stock_repository

import (
	"context"
	"errors"

	generic_repository "alves.com/modules/common/repo"
	stock_error "alves.com/modules/stock/errors"
	stock_model "alves.com/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByName(ctx context.Context, name string) (stock_model.StockEntity, error) {
	productGeneric, err := repo.Read(ctx, bson.M{"name": name})
	if err != nil {
		if errors.Is(err, generic_repository.ErrItemInexistent) {
			return stock_model.StockEntity{}, stock_error.ErrProductInexistent
		} else {
			return stock_model.StockEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	product, ok := productGeneric.(*stock_model.StockEntity)
	if !ok {
		return stock_model.StockEntity{}, stock_error.ErrCannotConvert
	}

	return *product, nil
}
