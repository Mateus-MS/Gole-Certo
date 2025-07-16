package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
	"github.com/google/uuid"
)

type service struct {
	repository repository.OrderRepository

	// Dependencies
	userServ userservice.Service
	prodServ productservice.Service
}

func New(repo repository.OrderRepository, userServ userservice.Service, prodServ productservice.Service) *service {
	return &service{
		repository: repo,
		userServ:   userServ,
		prodServ:   prodServ,
	}
}

func (s *service) Register(userID string, products []product.Product) (_ string, err error) {
	// 1 - Check if the received user exists
	if _, err = s.userServ.Search(userID); err != nil {
		return "", err
	}

	// 2 - Check if the received product list match existing products
	// NOTE: currently, it's not checking, it's using a mock, always returning true :P
	for _, product := range products {
		if _, err = s.prodServ.Search(product.ProductID); err != nil {
			return "", err
		}
	}

	// 3 - Create the structure to save in DB
	ord := order.New(
		userID,              // UserID
		uuid.New().String(), // OrderID
		"batching",          // State
		products,            // Products
	)

	// 4 - Save in DB
	if err = s.repository.Save(ord); err != nil {
		return "", err
	}

	// 5 - Check if the batching is full
	// NOTE: Whenever a new order is registered, it is registered as `batching` which means that we waiting till we have 1000 products
	// waiting to be ordered from `Duff Beer`. The point is, whenever we register a new order, we need to check if now, we have enough products
	// in the waiting list

	return ord.OrderID, nil
}

type SearchFilter struct {
	State  string
	UserID string
}

func (s *service) Search(filter SearchFilter) (ord order.Order, err error) {
	if filter.State != "" {
		println("Searching all orders with state: " + filter.State)
	}

	if filter.UserID != "" {
		println("Searching all orders of user: " + filter.UserID)
	}

	return ord, err
}
