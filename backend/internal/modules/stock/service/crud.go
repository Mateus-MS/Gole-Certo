package stock_service

import (
	"context"

	stock_model "alves.com/modules/stock/model"
)

func (s *service) Register(ctx context.Context, product stock_model.StockEntity) error {
	return s.repository.Create(ctx, &product)
}

func (s *service) ReadByName(ctx context.Context, name string) (stock_model.StockEntity, error) {
	return s.repository.ReadByName(ctx, name)
}
