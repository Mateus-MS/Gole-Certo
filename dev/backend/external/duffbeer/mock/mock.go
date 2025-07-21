package duffbeerService_mock

import (
	"errors"
	"math/rand"
	"time"

	duffbeer_service "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer"
	"github.com/google/uuid"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (c *service) SubmitOrder(ord duffbeer_service.Order) (resp duffbeer_service.OrderResponse, err error) {
	// Simulate API instability
	if rand.Intn(10) < 3 {
		return resp, errors.New("duff Beer is down (simulated)")
	}

	// Simulate processing time
	time.Sleep(700 * time.Millisecond)

	// Simulate OK response
	resp.OrderID = uuid.New().String()
	resp.Products = ord.Products

	return resp, nil
}
