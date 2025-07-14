package mock

import (
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

var TestError = errors.New("Test")

func (repo *ProductRepository) Search(identifier string) (prod product.Product, err error) {
	return prod, err
	// return prod, TestError
}
