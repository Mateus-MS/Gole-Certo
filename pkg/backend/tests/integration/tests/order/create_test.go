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

func TestOrderCreate_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	validUsername := "jhonDoe"
	validPassword := "jhonPass"
	validStockName := "Coca-cola"
	validStockQuantity := 30

	// insert user and stock
	userID := integration_fixtures.InsertTestUser(h.Ctx, h.DB.Database, validUsername, validPassword)
	stockID := integration_fixtures.InsertTestStock(h.Ctx, h.DB.Database, validStockName, validStockQuantity)

	// Create the order
	orderQuantity := int(float32(validStockQuantity) * 0.3)
	productsMap := map[primitive.ObjectID]int{stockID: orderQuantity}
	orderOBJ := order_model.New(userID, productsMap)

	// Try register the order
	require.NoError(t, h.Services.Order.Create(h.Ctx, *orderOBJ))

	// Assert
	entityStockUpdated, err := h.Services.Stock.ReadByID(h.Ctx, stockID)
	require.NoError(t, err)
	assert.Equal(t, validStockQuantity-orderQuantity, entityStockUpdated.Quantity)
}
