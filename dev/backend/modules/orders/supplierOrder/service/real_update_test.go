package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister_SuccessUpdatingProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered product
	prod := producttestutils.GetMockRegistered(t, app)[0]

	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	prodSum := order.Products[0].Quantity

	// It should create a new order batch
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Create a new order with the same product
	order, _ = supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	prodSum += order.Products[0].Quantity

	// It should update the existing order batch doubling the quantity
	ordID, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	ordOBJ, _ := primitive.ObjectIDFromHex(ordID)
	ordDB, err := app.Services.SupplierOrder.ReadByOrderID(ordOBJ)
	assert.NoError(t, err)
	assert.Equal(t, ordDB.Products[0].Quantity, prodSum)
}

func TestRegister_SuccessAppendingProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered product
	prod1 := producttestutils.GetMockRegistered(t, app)[0]

	// Register another product
	prod2, err := product.New(
		"Pepsi",
		"Pepsi",
		"1.99",
		50,
	)
	assert.NoError(t, err)
	assert.NoError(t, app.Services.Product.Create(prod2))

	order1, _ := supplierOrder.New(
		[]product.Product{prod1},
		"batching",
	)

	// It should create a new order batch
	ordID1, err := app.Services.SupplierOrder.Register(order1)
	assert.NoError(t, err)

	// Now creating a completely new order
	order2, _ := supplierOrder.New(
		[]product.Product{prod2},
		"batching",
	)

	// It should instade of create a new batch, merge it with the existing one
	ordID2, err := app.Services.SupplierOrder.Register(order2)
	assert.NoError(t, err)
	assert.Equal(t, ordID1, ordID2)
}
