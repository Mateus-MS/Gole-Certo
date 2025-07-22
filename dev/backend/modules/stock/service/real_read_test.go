package stock_service_test

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
)

func TestRead_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Try register into DB
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")

	// Search on DB
	prodDB, _ := app.Services.Product.ReadByName(prod.Name)

	// Assert
	assert.Equal(t, prod, prodDB)
}

func TestRead_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	_, err := app.Services.Product.ReadByName("Coca cola")
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
