package order_service_test

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/order"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
)

func TestCreate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a product registered
	prods := producttestutils.GetMockRegistered(t, app)

	ordr := ordertestutils.GetMock(prods)

	// Try register into DB
	_, err := app.Services.Order.Create(ordr)
	assert.NoError(t, err, "Expect no errors")
}

func TestCreate_InexistentProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a NON product registered
	prods := producttestutils.GetMock()

	ordr := ordertestutils.GetMock(prods)

	// Try register into DB
	_, err := app.Services.Order.Create(ordr)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
