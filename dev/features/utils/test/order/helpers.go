package ordertestutils

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	orderservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/order"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
)

// Try to register the received order into received application service
func Register(t *testing.T, app *testutils.Application, userID string, prods []product.Product) (ordID string, err error) {
	t.Helper()

	if ordID, err = app.Services.Order.Create(userID, prods); err != nil {
		return ordID, err
	}

	return ordID, nil
}

// Try to search with the received filter into the received application servie
func Search(t *testing.T, app *testutils.Application, filter orderservice.QueryFilter) (ord order.Order, err error) {
	t.Helper()

	if ord, err = app.Services.Order.Read(filter); err != nil {
		return ord, err
	}

	return ord, nil
}
