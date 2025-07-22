package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/order"
	"github.com/stretchr/testify/assert"
)

func TestRegister_SuccessUpdatingProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the first order
	order1, _, prod := ordertestutils.GetUnregisteredMock(t, app)
	prodOrderQuantity := prod.Quantity * 2

	// Register it
	// It should create a new order batch
	_, err := app.Services.SupplierOrder.Register(order1)
	assert.NoError(t, err)

	// Create a new order with the same product
	order2, _ := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{prod},
		"batching",
	)

	// try to register it again
	// It should sum the order products with the existing "batch"
	orderHEX, err := app.Services.SupplierOrder.Register(order2)
	assert.NoError(t, err)

	// To confirm
	// Search for the received ID
	ordDB, err := app.Services.SupplierOrder.ReadByOrderID(orderHEX)
	assert.NoError(t, err)

	// The total quantity of this order should be twice the prod, since we order it two times
	assert.Equal(t, prodOrderQuantity, ordDB.TotalQuantity)
}

func TestRegister_SuccessAppendingProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Register two different products into DB
	stock1, err := product.New(
		"Pepsi zero",
		"Pepsi",
		"1.99",
		50,
	)
	assert.NoError(t, err)
	assert.NoError(t, app.Services.Product.Create(stock1))

	stock2, err := product.New(
		"Coca cola zero",
		"Coca cola",
		"1.99",
		50,
	)
	assert.NoError(t, err)
	assert.NoError(t, app.Services.Product.Create(stock2))

	// Get the stock in supplier format
	prod1 := stock1.GetInSupplierFormat()
	prod1.Quantity = 200 // Ordered quantity

	prod2 := stock2.GetInSupplierFormat()
	prod2.Quantity = 400 // Ordered quantity

	// Create the first order with the first prod
	order1, _ := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{prod1},
		"batching",
	)

	// Order for the first order
	// It should create a new order "batch"
	ordID1, err := app.Services.SupplierOrder.Register(order1)
	assert.NoError(t, err)

	// Now a new order with a different product
	order2, _ := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{prod2},
		"batching",
	)

	// Order for the second order
	// It should append the order to the existing "batch"
	ordID2, err := app.Services.SupplierOrder.Register(order2)
	assert.NoError(t, err)

	// To check if the order was appended instead of a new one created
	// just check if the returned IDs of both register methods is equal
	assert.Equal(t, ordID1, ordID2)
}
