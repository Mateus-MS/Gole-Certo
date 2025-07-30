package costumerOrder_service

import (
	"errors"

	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	costumerOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository *costumerOrder_repository.Repository

	// Dependencies
	userService  contracts.User_Service
	stockService contracts.Stock_Service
}

var (
	ErrInsufficientStock = errors.New("the order is ordering more items than there is in stock")
)

func New(coll *mongo.Collection) service {
	return service{
		repository: costumerOrder_repository.New(coll),
	}
}

func (s *service) SetUserService(userService contracts.User_Service) {
	s.userService = userService
}

func (s *service) SetStockService(stockService contracts.Stock_Service) {
	s.stockService = stockService
}

func (s *service) Repo() *costumerOrder_repository.Repository {
	return s.repository
}
