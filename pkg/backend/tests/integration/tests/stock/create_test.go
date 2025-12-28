package integration_stock_test

import (
	"testing"

	stock_model "alves.com/backend/modules/stock/model"
	integration_setup "alves.com/backend/tests/integration/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStockCreate_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	var stock = *stock_model.New("Coca-cola", 20)

	// Try register
	{
		err := h.Services.Stock.Create(h.Ctx, stock)
		require.NoError(t, err, "stock creation should not return an error")
	}

	// Try to query it by ID
	{
		stockEntity, err := h.Services.Stock.ReadByID(h.Ctx, stock.ID)
		assert.NoError(t, err, "stock query should not return an error")
		assert.Equal(t, stock.Name, stockEntity.Name, "stock query should find a stock with the same name as queryied")
	}
}
