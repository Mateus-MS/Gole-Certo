package order_service

import (
	order "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/order/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	Create(order.Order) (string, error)
	Read(bson.M) (order.Order, error)
	ReadByOrderID(string) (order.Order, error)
}
