package order_service

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
)

func (s *service) Create(ctx context.Context, order order_model.OrderEntity) error {
	return s.repository.Create(ctx, order)
}
