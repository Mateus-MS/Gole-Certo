package order_service

import (
	"context"

	order_model "alves.com/backend/modules/order/model"
	order_repository "alves.com/backend/modules/order/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// DB crud functions to be exported
	Create(context.Context, order_model.OrderEntity) error
}

type service struct {
	repository *order_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: order_repository.New(coll),
	}
}
