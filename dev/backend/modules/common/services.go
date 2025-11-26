package contracts

import (
	"context"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	costumerOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/repository/mongo"
	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository/mongo"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	stock_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/repository/mongo"
	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	user_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/repository/mongo"
)

type SupplierOrder_Service interface {
	// Through the repository layer It will:
	// - Check all products received are valids
	// - Check if there is some batch... well, batching
	// - Will update the existing batch with the sended products
	// - If no batch, creates a new one
	Register(context.Context, supplierOrder.SupplierOrder) (string, error)

	SetStockService(Stock_Service)

	// Exposes the repository layer to the service layer
	Repo() *supplierOrder_repository.Repository
}

type CostumerOrder_Service interface {
	Register(context.Context, costumerOrder.CostumerOrder) (string, error)

	SetStockService(Stock_Service)
	SetUserService(User_Service)

	// Exposes the repository layer to the service layer
	Repo() *costumerOrder_repository.Repository
}

type Stock_Service interface {
	Register(context.Context, product.ProductStock) error

	DeductFromStock(context.Context, product.ProductStock, int64) error

	// Service setters
	SetSupplierOrderService(SupplierOrder_Service)

	// Exposes the repository layer to the service layer
	Repo() *stock_repository.Repository
}

type User_Service interface {
	Register(context.Context, user.User) error

	// Exposes the repository layer to the service layer
	Repo() *user_repository.Repository
}
