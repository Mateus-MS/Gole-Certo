package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/order"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
)

func TestCreate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the new supplier order OBJ
	order, _, _ := ordertestutils.GetUnregisteredMock(t, app)

	// Try to save it on DB
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)
}

func TestCreate_NotRegisteredProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a NOT registered product
	stock := producttestutils.GetMock()[0]

	// Get the NON registered product in SupplierProduct format
	prod := stock.GetInSupplierFormat()
	prod.Quantity = 200

	// Create the new supplier order OBJ
	order, _ := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{prod},
		"batching",
	)

	// Try to save it on DB
	_, err := app.Services.SupplierOrder.Register(order)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
