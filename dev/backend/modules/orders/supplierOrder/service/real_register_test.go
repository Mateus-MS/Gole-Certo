package supplierOrder_service_test

import (
	"testing"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
)

func TestCreate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered product
	prod := producttestutils.GetMockRegistered(t, app)[0]

	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)

	// Try to save it on DB
	_, err := app.Services.SupplierOrder.Register(order)
	assert.NoError(t, err)
}

func TestCreate_NotRegisteredProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a NOT registered product
	prod := producttestutils.GetMock()[0]

	order, _ := supplierOrder.New(
		[]product.Product{prod},
		"batching",
	)

	// Try to save it on DB
	_, err := app.Services.SupplierOrder.Register(order)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
