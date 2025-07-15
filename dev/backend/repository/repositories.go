package repository

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
)

type Repositories struct {
	User interface {
		Save(cli user.User) (err error)
		Search(identifier string) (c user.User, err error)
	}
	Product interface {
		Search(identifier string) (p product.Product, err error)
	}
	Order interface {
		Save(ord order.Order) (err error)
	}
}
