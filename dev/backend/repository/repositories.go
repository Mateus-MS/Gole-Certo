package repository

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
)

// TODO: Seems strange to inittiate the interfaces here

type UserRepository interface {
	Save(cli user.User) (err error)
	Search(identifier string) (c user.User, err error)
}

type ProductRepository interface {
	Search(identifier string) (p product.Product, err error)
}

type OrderRepository interface {
	Save(ord order.Order) (err error)
}

type Repositories struct {
	User    UserRepository
	Product ProductRepository
	Order   OrderRepository
}
