package stock_service

import (
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Create(product.ProductStock) error

	// ReadManyAfterID(bson.M, string, int64) ([]Product, error)
	ReadByID(string) (product.ProductStock, error)
	ReadByName(string) (product.ProductStock, error)

	UpdateByID(product.ProductStock) error

	DeleteByID(string) error
	DeleteByName(string) error

	// Utils
	ValidateProductByID(string) bool

	// Base functions
	Delete(bson.M) error
	Read(bson.M) (product.ProductStock, error)

	// TODO
	// ApplyStockReduction(string, int) error
}

// TODO: ReadMany
