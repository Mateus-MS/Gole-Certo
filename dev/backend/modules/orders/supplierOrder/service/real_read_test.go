package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordersuppliertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/orders/supplier"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReadByState_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the new supplier order OBJ
	order, _, _ := ordersuppliertestutils.GetUnregisteredMock(t, app)

	// Register the supplier order into DB
	// This should create a new one
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	orderDB, err := app.Services.SupplierOrder.ReadOneByState("batching")
	assert.NoError(t, err)
	assert.Equal(t, order.ID, orderDB.ID)
}

func TestReadByState_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	// Search for it
	_, err := app.Services.SupplierOrder.ReadOneByState("batching")
	assert.ErrorIs(t, err, supplierOrder_repository.ErrOrderNotFound)
}

func TestReadByOrderID_Success(t *testing.T) {
	app := testutils.SetupTest(t)

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

	// Register the supplier order into DB
	// This should create a new one
	ordID, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	orderDB, err := app.Services.SupplierOrder.ReadByOrderID(ordID)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, orderDB.ID)
}

func TestReadByOrderID_InvalidIDType(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create a product into DB
	prod := producttestutils.GetMockRegistered(t, app)[0]

	// Try Search for a invalid object
	// It should accept only string or primitive.ObjectID
	_, err := app.Services.SupplierOrder.ReadByOrderID(prod)
	assert.ErrorIs(t, err, utils.ErrInvalidIDType)
}

func TestReadByOrderID_InvalidStringFormat(t *testing.T) {
	app := testutils.SetupTest(t)

	// Try search for a invalid primitive.ObjectID format
	_, err := app.Services.SupplierOrder.ReadByOrderID("123918273")
	assert.ErrorIs(t, err, utils.ErrInvalidIDFormat)
}

func TestReadByOrderID_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	// A inexistent ID on DB
	testID, _ := primitive.ObjectIDFromHex("687e127b86885f3fda9ea926")

	// Try search for it
	_, err := app.Services.SupplierOrder.ReadByOrderID(testID)
	assert.ErrorIs(t, err, supplierOrder_repository.ErrOrderNotFound)
}
