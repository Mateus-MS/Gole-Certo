package duffbeerService

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Order(ord duffbeer.Order) (resp duffbeer.OrderResponse) {
	if resp, err := app.GetInstance().Duffbeer.SubmitOrder(ord); err != nil {
		return resp
	}
	return resp
}
