package product_service

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Create(product.Product) error

	Read(bson.M) (product.Product, error)
	// ReadManyAfterID(bson.M, string, int64) ([]product.Product, error)
	ReadByID(string) (product.Product, error)
	ReadByName(string) (product.Product, error)

	UpdateByID(product.Product) error

	Delete(bson.M) error
	DeleteByID(string) error
	DeleteByName(string) error

	// Utils
	ValidateList([]product.Product) bool
}

// TODO: ReadMany
