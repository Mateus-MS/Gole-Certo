package order_repository

import (
	"context"
	"errors"

	order "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/order/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrOrderNotFound = errors.New("order does not exists in DB")
)

type Repository struct {
	Collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{Collection: coll}
}

func (repo *Repository) Create(ord order.Order) (err error) {
	if _, err = repo.Collection.InsertOne(context.TODO(), ord); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Read(queryFilter bson.M) (ord order.Order, err error) {
	if err = repo.Collection.FindOne(context.TODO(), queryFilter).Decode(&ord); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ord, ErrOrderNotFound
		}
	}

	return ord, nil
}
