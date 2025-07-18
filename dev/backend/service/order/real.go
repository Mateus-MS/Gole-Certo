package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *service) Create(userID string, products []product.Product) (_ string, err error) {
	// 1 - Check if the received user exists
	if _, err = s.userServ.Read(userID); err != nil {
		return "", err
	}

	// 2 - Check if the received product list match existing products
	// NOTE: currently, it's not checking, it's using a mock, always returning true :P
	for _, product := range products {

		_, err = s.prodServ.Read(
			productservice.QueryFilter{
				ID: product.ProductID,
			},
		)

		if err != nil {
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
	if err = s.repository.Create(ord); err != nil {
		return "", err
	}

	// 5 - Check if the batching is full
	// NOTE: Whenever a new order is registered, it is registered as `batching` which means that we waiting till we have 1000 products
	// waiting to be ordered from `Duff Beer`. The point is, whenever we register a new order, we need to check if now, we have enough products
	// in the waiting list

	return ord.OrderID, nil
}

type QueryFilter struct {
	State   string
	UserID  string
	OrderID string
}

func (s *service) Read(filter QueryFilter) (ord order.Order, err error) {
	queryFilter := bson.M{}

	// Dinamically build the filter
	if filter.State != "" {
		queryFilter["state"] = filter.State
	}

	if filter.UserID != "" {
		queryFilter["userID"] = filter.UserID
	}

	if filter.OrderID != "" {
		queryFilter["_id"] = filter.OrderID
	}

	// Perform the query
	if ord, err = s.repository.Read(queryFilter); err != nil {
		return ord, err
	}

	return ord, nil
}
