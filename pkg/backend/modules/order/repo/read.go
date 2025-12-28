package order_repository

import (
	"context"

	order_error "alves.com/backend/modules/order/errors"
	order_model "alves.com/backend/modules/order/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *Repository) ReadAllByUserID(ctx context.Context, userID primitive.ObjectID) ([]order_model.OrderEntity, error) {
	cursor, err := repo.Collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []order_model.OrderEntity
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, order_error.ErrOrderInexistent
	}

	return orders, nil
}
