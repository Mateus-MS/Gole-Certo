package order_repository

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
	user_error "alves.com/backend/modules/user/errors"
)

func (repo *Repository) Create(ctx context.Context, orderEntity order_model.OrderEntity) error {
	_, err := repo.Collection.InsertOne(ctx, orderEntity)

	if err != nil {
		return user_error.ErrUserAlreadyExists
	}

	return nil
}
