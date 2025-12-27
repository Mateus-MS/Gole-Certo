package stock_service

import (
	"context"

	stock_model "alves.com/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) Create(ctx context.Context, stock stock_model.StockEntity) error {
	return s.repository.Create(ctx, stock)
}
func (s *service) ReadByName(ctx context.Context, name string) (*stock_model.StockEntity, error) {
	return s.repository.ReadByName(ctx, name)
}
func (s *service) ReadByID(ctx context.Context, id primitive.ObjectID) (*stock_model.StockEntity, error) {
	return s.repository.ReadByID(ctx, id)
}
func (s *service) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return s.repository.DeleteByID(ctx, id)
}
