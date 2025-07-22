package stock_service_test

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Try register into DB
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")

	// Get the product snapshot
	prodSnapshot, err := app.Services.Product.ReadByName(prod.Name)
	assert.NoError(t, err)

	// Check if what is in DB is equals to what was sended
	assert.Equal(t, prod, prodSnapshot)

	// Change the price of the product
	prodUpdate := prodSnapshot
	prodUpdate.Price, err = primitive.ParseDecimal128("4.99")
	assert.NoError(t, err)

	// Try update it
	err = app.Services.Product.UpdateByID(prodUpdate)
	assert.NoError(t, err)

	// Search for the new state of the product
	prodChanged, err := app.Services.Product.ReadByName(prodSnapshot.Name)
	assert.NoError(t, err)

	assert.Equal(t, prodUpdate, prodChanged)
}

func TestUpdate_InvalidPrice(t *testing.T) {
	app := testutils.SetupTest(t)
	prod := producttestutils.GetMock()[0]

	// register
	app.Services.Product.Create(prod)

	// Invalidate the price
	prod.Price, _ = primitive.ParseDecimal128("-20")

	// Try update it
	err := app.Services.Product.UpdateByID(prod)
	assert.ErrorIs(t, err, product.ErrInvalidPrice)
}

func TestUpdate_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	err := app.Services.Product.UpdateByID(prod)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
