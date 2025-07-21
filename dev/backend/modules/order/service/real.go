package order_service

import (
	order "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/order/model"
	order_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/order/repository"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	product_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/service"
	user_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository order_repository.Repository

	// Dependencies
	userServ user_service.Service
	prodServ product_service.Service
}

func New(coll *mongo.Collection, userServ user_service.Service, prodServ product_service.Service) *service {
	return &service{
		repository: *order_repository.New(coll),
		userServ:   userServ,
		prodServ:   prodServ,
	}
}

func (s *service) Create(ord order.Order) (_ string, err error) {
	// 1 - Check if the received user exists
	// if _, err = s.userServ.Read(ord.UserIdentifier); err != nil {
	// 	return "", err
	// }

	// 2 - Check if the received product list match existing products
	for _, prod := range ord.Product {

		_, err = s.prodServ.ReadByName(prod.Name)

		if err != nil {
			return "", product.ErrProductInexistent
		}
	}

	// 3 - Save in DB
	if err = s.repository.Create(ord); err != nil {
		return "", err
	}

	// 4 - Check if the batching is full
	// NOTE: Whenever a new order is registered, it is registered as `batching` which means that we waiting till we have 1000 products
	// waiting to be ordered from `Duff Beer`. The point is, whenever we register a new order, we need to check if now, we have enough products
	// in the waiting list

	return ord.OrderID.Hex(), nil
}

func (s *service) Read(filter bson.M) (ord order.Order, err error) {
	if ord, err = s.repository.Read(filter); err != nil {
		return ord, err
	}

	return ord, nil
}
func (s *service) ReadByOrderID(ordID string) (ord order.Order, err error) {
	if ord, err = s.repository.Read(bson.M{"_id": ordID}); err != nil {
		return ord, err
	}

	return ord, nil
}
