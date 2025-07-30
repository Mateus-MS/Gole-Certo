package supplierOrder_service

import (
	"errors"

	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrOrderStateMustBeBatching = errors.New("supplier order status must be batching")
)

type service struct {
	repository *supplierOrder_repository.Repository

	// Dependencies
	stockService contracts.Stock_Service
}

func (s *service) SetStockService(stockService contracts.Stock_Service) {
	s.stockService = stockService
}

func New(coll *mongo.Collection) service {
	return service{
		repository: supplierOrder_repository.New(coll),
	}
}

func (s *service) Repo() *supplierOrder_repository.Repository {
	return s.repository
}
