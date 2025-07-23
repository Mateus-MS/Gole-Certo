package contracts

import (
	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

type SupplierOrder_Service interface {
	// Through the repository layer It will:
	// - Check all products received are valids
	// - Check if there is some batch... well, batching
	// - Will update the existing batch with the sended products
	// - If no batch, creates a new one
	Register(supplierOrder.SupplierOrder) (string, error)

	ReadByOrderID(any) (supplierOrder.SupplierOrder, error)
	ReadOneByState(string) (supplierOrder.SupplierOrder, error)
	ReadManyByState(string, int) ([]supplierOrder.SupplierOrder, error)

	UpdateByID(supplierOrder.SupplierOrder) error

	SetStockService(Stock_Service)
}

type CostumerOrder_Service interface {
	Register(costumerOrder.CostumerOrder) (string, error)

	// C R U D
	Create(costumerOrder.CostumerOrder) (string, error)
}

type Stock_Service interface {
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
	ApplyStockReduction(string, int64) error

	// Service setters
	SetSupplierOrderService(SupplierOrder_Service)
}

type User_Service interface {
	Create(user.User) error
	Read(string) (user.User, error)
}
