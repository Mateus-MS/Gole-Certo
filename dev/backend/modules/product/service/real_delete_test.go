package product_service_test

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Register
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")

	// Delete
	err = app.Services.Product.DeleteByName(prod.Name)
	assert.NoError(t, err)

	// To confirm that it was deleted from DB, try to search for it
	prodDB, err := app.Services.Product.ReadByName(prod.Name)
	assert.Empty(t, prodDB)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}

func TestDeleteByName_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Register
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")

	// Delete
	err = app.Services.Product.DeleteByName(prod.Name)
	assert.NoError(t, err)

	// To confirm that it was deleted from DB, try to search for it
	prodDB, err := app.Services.Product.ReadByName(prod.Name)
	assert.Empty(t, prodDB)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}

func TestDeleteByID_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Register
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")

	// Query by the name just to get the ID generated
	prodDB, _ := app.Services.Product.ReadByName(prod.Name)

	// Delete by the ID
	err = app.Services.Product.DeleteByID(prodDB.ProductID.Hex())
	assert.NoError(t, err)

	// To confirm that it was deleted from DB, try to search for it
	assert.NoError(t, err)
	prodDB, err = app.Services.Product.ReadByName(prod.Name)
	assert.Empty(t, prodDB)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}

func TestDelete_Inexistent(t *testing.T) {
	app := testutils.SetupTest(t)

	// Try delete the product
	err := app.Services.Product.DeleteByName("Coca cola")
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}
