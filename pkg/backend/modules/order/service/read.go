package order_service

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) ReadAllByUserID(ctx context.Context, userID primitive.ObjectID) ([]order_model.OrderEntity, error) {
	_, err := s.user_service.ReadByID(ctx, userID)
	if err != nil {
		return []order_model.OrderEntity{}, err
	}

	return s.repository.ReadAllByUserID(ctx, userID)
}
