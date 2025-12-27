package order_service

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
)

func (s *service) Create(ctx context.Context, order order_model.OrderEntity) error {
	// Check if the given user exists
	_, err := s.user_service.ReadByID(ctx, order.UserID)
	if err != nil {
		return err
	}

	// Check if the given products exists
	for prodId := range order.Products {
		_, err := s.stock_service.ReadByID(ctx, prodId)
		if err != nil {
			return err
		}
	}

	// Update the stock quantity
	for prodId, quantity := range order.Products {
		err := s.stock_service.AtomicDecreaseStockByID(ctx, prodId, quantity)
		if err != nil {
			return err
		}
	}

	return s.repository.Create(ctx, order)
}
