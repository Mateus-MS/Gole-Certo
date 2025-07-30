package stock_repository

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Alias
type Product = product.ProductStock

type Repository struct {
	collection *mongo.Collection
}

func New(coll *mongo.Collection) *Repository {
	return &Repository{collection: coll}
}
