package duffbeer_service

import product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"

type Service interface {
	SubmitOrder(Order) (OrderResponse, error)
}

// Order contains the reseller information `Gole Certo` and the products being ordered.
// These fields like CNPJ and Address are static and can be loaded from .env
type Order struct {
	cnpj     string
	address  string
	Products []product.Product
}

func NewOrder(cnpj, address string, prods []product.Product) Order {
	return Order{
		cnpj:     cnpj,
		address:  address,
		Products: prods,
	}
}

type OrderResponse struct {
	OrderID  string
	Products []product.Product
}
