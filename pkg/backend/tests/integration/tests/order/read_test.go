package integration_order_test

import (
	"testing"

	order_model "alves.com/backend/modules/order/model"
	integration_setup "alves.com/backend/tests/integration/setup"
	integration_fixtures "alves.com/backend/tests/integration/setup/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderReadAllOrdersFromUser_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	// Arrange: create user
	username := "jhonDoe_" + t.Name()
	password := "jhonPass"
	userID := integration_fixtures.InsertTestUser(h.Ctx, h.DB.Database, username, password)

	// Arrange: stock data
	stockData := map[string]int{
		"Coca-cola":     20,
		"Pepsi":         50,
		"Guaraná":       43,
		"Fanta Laranja": 14,
	}

	// Arrange + Act: create one order per product
	for name, quantity := range stockData {
		stockID := integration_fixtures.InsertTestStock(h.Ctx, h.DB.Database, name, quantity)
		orderQuantity := quantity - int(float32(quantity)*0.2)
		order := order_model.New(userID, map[primitive.ObjectID]int{stockID: orderQuantity})
		require.NoError(t, h.Services.Order.Create(h.Ctx, *order))
	}

	// Act: read all orders
	orders, err := h.Services.Order.ReadAllByUserID(h.Ctx, userID)
	require.NoError(t, err)

	// Assert: number of orders matches products
	assert.Equal(t, len(stockData), len(orders))
}
