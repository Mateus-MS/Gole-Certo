package costumerOrder_service

import (
	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	costumerOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/repository"
	product_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/service"
	user_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/service"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository costumerOrder_repository.Repository

	// Dependencies
	userService user_service.Service
	prodService product_service.Service
}

func New(coll *mongo.Collection, userService user_service.Service, prodService product_service.Service) service {
	return service{
		repository:  *costumerOrder_repository.New(coll),
		userService: userService,
		prodService: prodService,
	}
}

func (s *service) Register(ord costumerOrder.CostumerOrder) (ordID string, err error) {
	// check if the received products list really exists on DB
	for _, prod := range ord.Products {
		if _, err := s.prodService.ReadByID(prod.GetProductID()); err != nil {
			return "", err
		}
	}

	// Check if the received user exists on DB
	if _, err = s.userService.Read(ord.UserID); err != nil {
		return "", err
	}

	// Create the order
	return s.Create(ord)
}

// C R U D

func (s *service) Create(ord costumerOrder.CostumerOrder) (ordID string, err error) {
	if err = s.repository.Create(ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}
