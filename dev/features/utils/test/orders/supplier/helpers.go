package ordersuppliertestutils

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
)

func GetUnregisteredMock(t *testing.T, app *testutils.Application) (supplierOrder.SupplierOrder, product.ProductStock, *supplierOrder.SupplierProduct) {
	// Create a product into DB
	stock := producttestutils.GetMockRegistered(t, app)[0]

	// Get the registered product in SupplierProduct format
	prod := stock.GetInSupplierFormat()
	prod.Quantity = 100

	// Create the new supplier order OBJ
	order, _ := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{prod},
		"batching",
	)

	return order, stock, prod
}
