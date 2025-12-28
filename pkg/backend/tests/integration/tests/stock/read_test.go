package integration_stock_test

import (
	"testing"

	integration_setup "alves.com/backend/tests/integration/setup"
	integration_fixtures "alves.com/backend/tests/integration/setup/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStockRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	var validStockName = "Coca-cola"
	var validStockQuantity = 30

	// Direct insert the user into DB
	stockID := integration_fixtures.InsertTestStock(h.Ctx, h.DB.Database, validStockName, validStockQuantity)

	// Try to read the user
	{
		stockEntity, err := h.Services.Stock.ReadByID(h.Ctx, stockID)
		require.NoError(t, err, "stock query should not return an error")
		assert.Equal(t, validStockName, stockEntity.Name, "stock ID query should return a stock entity with the same name as queryied")
	}
}
