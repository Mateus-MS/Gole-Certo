package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReadByState_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create a order into DB
	prod := producttestutils.GetMockRegistered(t, app)[0]
	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	orderDB, err := app.Services.SupplierOrder.ReadOneByState("batching")
	assert.NoError(t, err)
	assert.Equal(t, order, orderDB)
}

func TestReadByState_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	// Search for it
	_, err := app.Services.SupplierOrder.ReadOneByState("batching")
	assert.ErrorIs(t, err, supplierOrder_repository.ErrOrderNotFound)
}

func TestReadByOrderID_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create a new order
	// Create a order into DB
	prod := producttestutils.GetMockRegistered(t, app)[0]
	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	ordID, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	orderDB, err := app.Services.SupplierOrder.ReadByOrderID(ordID)
	assert.NoError(t, err)
	assert.Equal(t, order, orderDB)
}

func TestReadByOrderID_InvalidIDType(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create a new order
	// Create a order into DB
	prod := producttestutils.GetMockRegistered(t, app)[0]
	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	_, err = app.Services.SupplierOrder.ReadByOrderID(prod)
	assert.ErrorIs(t, err, utils.ErrInvalidIDType)
}

func TestReadByOrderID_InvalidStringFormat(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create a order into DB
	prod := producttestutils.GetMockRegistered(t, app)[0]
	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)

	// Search for it
	_, err = app.Services.SupplierOrder.ReadByOrderID("123918273")
	assert.ErrorIs(t, err, utils.ErrInvalidIDFormat)
}

func TestReadByOrderID_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	testID, _ := primitive.ObjectIDFromHex("687e127b86885f3fda9ea926")

	// Search for it
	_, err := app.Services.SupplierOrder.ReadByOrderID(testID)
	assert.ErrorIs(t, err, supplierOrder_repository.ErrOrderNotFound)
}
