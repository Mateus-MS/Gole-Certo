package stock_service

import (
	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	stock_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Alias
type Stock = product.ProductStock

type service struct {
	repository *stock_repository.Repository

	supplierOrder contracts.SupplierOrder_Service
}

func (s *service) SetSupplierOrderService(supplierOrder contracts.SupplierOrder_Service) {
	s.supplierOrder = supplierOrder
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: stock_repository.New(coll),
	}
}

func (s *service) Repo() *stock_repository.Repository {
	return s.repository
}
