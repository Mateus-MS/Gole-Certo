package costumerOrder_service

import (
	"errors"

	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	costumerOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository costumerOrder_repository.Repository

	// Dependencies
	userService  contracts.User_Service
	stockService contracts.Stock_Service
}

var (
	ErrInsufficientStock = errors.New("the order is ordering more items than there is in stock")
)

func New(coll *mongo.Collection) service {
	return service{
		repository: *costumerOrder_repository.New(coll),
	}
}

func (s *service) SetUserService(userService contracts.User_Service) {
	s.userService = userService
}

func (s *service) SetStockService(stockService contracts.Stock_Service) {
	s.stockService = stockService
}

func (s *service) Register(ord costumerOrder.CostumerOrder) (ordID string, err error) {
	// Check if the received user exists on DB
	if _, err = s.userService.Read(ord.UserID); err != nil {
		return "", err
	}

	// check if the received products list really exists on DB
	for _, prod := range ord.Products {
		stock, err := s.stockService.ReadByID(prod.GetProductID())
		if err != nil {
			return "", err
		}

		// perhaps the desirable behavior is not this
		// maybe be flatten the order
		// or don't do nothing allowing this to continue normally
		if stock.Stock < prod.GetAmmount() {
			return "", ErrInsufficientStock
		}
	}

	// Create the order
	return s.Create(ord)

	// Remove from STOCK what was ordered
}

// C R U D

func (s *service) Create(ord costumerOrder.CostumerOrder) (ordID string, err error) {
	if err = s.repository.Create(ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}
