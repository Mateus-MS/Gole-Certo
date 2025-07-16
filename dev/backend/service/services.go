package service

import (
	duffbeerService "github.com/Mateus-MS/Gole-Certo/dev/backend/service/external/duffbeer"
	orderservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/order"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
)

type Services struct {
	User     userservice.Service
	Product  productservice.Service
	Order    orderservice.Service
	DuffBeer duffbeerService.Service
}
