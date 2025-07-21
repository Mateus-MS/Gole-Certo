package costumerOrder_repository

import (
	"context"
	"errors"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrOrderNotFound = errors.New("costumer order does not exists in db")
)

// Alias
type Order = costumerOrder.CostumerOrder

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{Collection: coll}
}

func (repo *Repository) Create(ord Order) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Read(filter bson.M) (ord Order, err error) {
	if err = repo.Collection.FindOne(context.TODO(), filter).Decode(&ord); err != nil {
		return ord, ErrOrderNotFound
	}

	return ord, nil
}
