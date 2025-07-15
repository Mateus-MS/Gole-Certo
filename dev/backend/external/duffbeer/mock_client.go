package duffbeer

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type MockClient struct {
}

func (c *MockClient) SubmitOrder(ord Order) (resp OrderResponse, err error) {
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
