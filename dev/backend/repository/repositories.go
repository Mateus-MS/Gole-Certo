package repository

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
)

// TODO: Seems strange to inittiate the interfaces here

type ClientRepository interface {
	Save(cli client.Client) (err error)
	Search(identifier string) (c client.Client, err error)
}

type ProductRepository interface {
	Search(identifier string) (p product.Product, err error)
}

type OrderRepository interface {
	Save(ord order.Order) (err error)
}

type Repositories struct {
	Client  ClientRepository
	Product ProductRepository
	Order   OrderRepository
}
