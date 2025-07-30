package costumerOrder_repository

import (
	"errors"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrOrderNotFound = errors.New("costumer order does not exists in db")
)

// Alias
type Order = costumerOrder.CostumerOrder

type Repository struct {
	collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{collection: coll}
}
