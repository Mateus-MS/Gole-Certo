package product_service_test

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	prod := producttestutils.GetMock()[0]

	// Try register into DB
	err := app.Services.Product.Create(prod)
	assert.NoError(t, err, "Expect no errors")
}

func TestCreate_NegativePrice(t *testing.T) {
	app := testutils.SetupTest(t)
	prod := producttestutils.GetMock()[0]

	invalidPrice, _ := primitive.ParseDecimal128("-2.99")
	prod.Price = invalidPrice

	// Try register the invalid product into DB
	err := app.Services.Product.Create(prod)
	assert.ErrorIs(t, err, product.ErrInvalidPrice)
}

func TestCreate_ZeroPrice(t *testing.T) {
	app := testutils.SetupTest(t)
	prod := producttestutils.GetMock()[0]

	invalidPrice, _ := primitive.ParseDecimal128("0")
	prod.Price = invalidPrice

	// Try register the invalid product into DB
	err := app.Services.Product.Create(prod)
	assert.ErrorIs(t, err, product.ErrInvalidPrice)
}

func TestCreate_NegativeStock(t *testing.T) {
	app := testutils.SetupTest(t)
	prod := producttestutils.GetMock()[0]

	prod.Quantity = -10

	// Try register the invalid product into DB
	err := app.Services.Product.Create(prod)
	assert.ErrorIs(t, err, product.ErrInvalidQuantity)
}

func TestCreate_EmptyName(t *testing.T) {
	app := testutils.SetupTest(t)
	prod := producttestutils.GetMock()[0]

	prod.Name = ""

	// Try register the invalid product into DB
	err := app.Services.Product.Create(prod)
	assert.ErrorIs(t, err, product.ErrInvalidName)
}
