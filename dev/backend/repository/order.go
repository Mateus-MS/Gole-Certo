package repository

import (
	"context"
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrOrderNotFound = errors.New("order does not exists in DB")
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func (repo *OrderRepository) Create(ord order.Order) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}

func (repo *OrderRepository) Read(queryFilter bson.M) (ord order.Order, err error) {
	if err = repo.Collection.FindOne(context.TODO(), queryFilter).Decode(&ord); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ord, ErrOrderNotFound
		}
	}

	return ord, nil
}
