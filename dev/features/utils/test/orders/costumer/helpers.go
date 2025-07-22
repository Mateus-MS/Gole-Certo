package ordercostumertestutils

import (
	"testing"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
)

func GetUnregisteredMock(t *testing.T, app *testutils.Application) (costumerOrder.CostumerOrder, product.ProductStock, *costumerOrder.CostumerProduct) {
	// Create a product into DB
	stock := producttestutils.GetMockRegistered(t, app)[0]

	// Get the registered product in SupplierProduct format
	prod := stock.GetInCostumerFormat()
	prod.Quantity = 100

	// Create the new supplier order OBJ
	order, _ := costumerOrder.New(
		[]costumerOrder.CostumerProduct{*prod},
		"batching",
	)

	return order, stock, prod
}
