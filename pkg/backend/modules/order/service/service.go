package order_service

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
	order_repository "alves.com/backend/modules/order/repo"
	stock_service "alves.com/backend/modules/stock/service"
	user_service "alves.com/backend/modules/user/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	Create(context.Context, order_model.OrderEntity) error

	ReadAllByUserID(context.Context, primitive.ObjectID) ([]order_model.OrderEntity, error)
}

type service struct {
	repository *order_repository.Repository

	user_service  user_service.IService
	stock_service stock_service.IService
}

func New(coll *mongo.Collection, userService user_service.IService, stockService stock_service.IService) *service {
	return &service{
		repository:    order_repository.New(coll),
		user_service:  userService,
		stock_service: stockService,
	}
}
