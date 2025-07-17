package orderservice_test

import (
	"testing"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	orderservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/order"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/order"
	usertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the user to register the order
	usr := usertestutils.CreateValidIndividual(t)

	// Register the user on mockDB
	usertestutils.Register(t, app, &usr)

	// Create the products of the order
	products := []product.Product{
		{
			ProductID: "123123123",
			Quantity:  25,
		},
	}

	// Register the order
	orderID, err := ordertestutils.Register(t, app, usr.GetIdentifier(), products)
	assert.NoError(t, err, "error while registering order in DB")
	assert.NotEmpty(t, orderID, "orderID should not be an empty string")

	// Test if the Order really is on DB
	ord, _ := ordertestutils.Search(t, app, orderservice.QueryFilter{OrderID: orderID})
	assert.EqualValues(t, orderID, ord.OrderID, "order ID from query should match the one returned on creation")
	assert.EqualValues(t, "033.355.662-38", ord.UserIdentifier, "user identifier from query should match the one used during order creation")
	assert.Equal(t, products, ord.Product, "products from query should match the ones used during order creation")
}

func TestCreateOrderWithoutRegisteredUser(t *testing.T) {
	app := testutils.SetupTest(t)

	// Create the products of the order
	products := []product.Product{
		{
			ProductID: "123123123",
			Quantity:  25,
		},
	}

	// Register the order
	orderID, err := ordertestutils.Register(t, app, "033.355.662-38", products)
	assert.ErrorIs(t, err, repository.ErrUserNotFound, "Expected User Not Found Error")
	assert.Empty(t, orderID, "OrderID should be empty")

	// Test if the Order is NOT on DB
	ord, err := ordertestutils.Search(t, app, orderservice.QueryFilter{OrderID: orderID})
	assert.ErrorIs(t, repository.ErrOrderNotFound, err, "since the order was not created, is should not be in DB")

	assert.EqualValues(t, orderID, ord.OrderID, "order ID from query should match the one returned on creation")
	assert.EqualValues(t, "", ord.UserIdentifier, "user identifier from query shouldn't exist")
	assert.Empty(t, ord.Product, "products from query should be empty")
}

// TODO: Create more tests when products are implemented, like:
// - what happens if receive a list with the duplicated product
// - test the quantity, shouldn't be 0
// - what happens if receive two identical orders from the same user in a short period of time
// - ...
