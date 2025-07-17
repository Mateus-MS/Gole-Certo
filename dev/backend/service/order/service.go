package orderservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
)

type Service interface {
	Register(string, []product.Product) (string, error)
	Search(QueryFilter) (order.Order, error)
}
