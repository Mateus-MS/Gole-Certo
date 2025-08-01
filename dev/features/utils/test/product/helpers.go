package producttestutils

import (
	"testing"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
)

func GetMock() []product.ProductStock {
	prod, _ := product.New(
		"Super BOCK Black",
		"Super BOCK",
		"1.99",
		50,
	)

	return []product.ProductStock{prod}
}

func GetMockRegistered(t *testing.T, app *testutils.Application) []product.ProductStock {
	t.Helper()

	prod, _ := product.New(
		"Super BOCK Black",
		"Super BOCK",
		"1.99",
		50,
	)

	app.Services.Stock.Create(prod)

	return []product.ProductStock{prod}
}
