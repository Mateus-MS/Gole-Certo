package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
)

type Service interface {
	Create(string, []product.Product) (string, error)
	Read(QueryFilter) (order.Order, error)
}
