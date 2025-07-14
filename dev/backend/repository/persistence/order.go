package persistence

import (
	"context"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	DB *mongo.Client
}

func (rep *OrderRepository) Save(ord order.Order) (err error) {
	collection := rep.DB.Database("goleCertoDB").Collection("orders")

	if _, err = collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}
