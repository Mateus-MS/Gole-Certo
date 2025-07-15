package duffbeer

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
)

type Client interface {
	SubmitOrder(Order) (OrderResponse, error)
}

// Order contains the reseller information `Gole Certo` and the products being ordered.
// These fields like CNPJ and Address are static and can be loaded from .env
type Order struct {
	CNPJ     string
	Address  string
	Products []product.Product
}

type OrderResponse struct {
	OrderID  string
	Products []product.Product
}

func NewOrder(cnpj, address string, prods []product.Product) Order {
	return Order{
		CNPJ:     cnpj,
		Address:  address,
		Products: prods,
	}
}
