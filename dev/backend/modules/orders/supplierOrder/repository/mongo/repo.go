package supplierOrder_repository

import (
	"errors"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Errors
var (
	ErrOrderNotFound = errors.New("supplier order does not exists in db")
)

// Alias
type Order = supplierOrder.SupplierOrder

type Repository struct {
	collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{collection: coll}
}
