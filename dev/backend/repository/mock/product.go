package mock

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (repo *ProductRepository) Search(identifier string) (prod product.Product, err error) {
	return prod, err
}
