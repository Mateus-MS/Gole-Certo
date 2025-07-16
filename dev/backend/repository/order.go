package repository

import (
	"context"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func (repo *OrderRepository) Save(ord order.Order) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}
