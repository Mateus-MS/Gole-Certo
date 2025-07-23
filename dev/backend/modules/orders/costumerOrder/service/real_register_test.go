package costumerOrder_service_test

import (
	"testing"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	costumerOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/service"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	user_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/repository"
	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordercostumertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/orders/costumer"
	producttestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/product"
	usertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered user in DB
	usr := usertestutils.GetMockRegistered(t, app)

	// Get a registered product in DB
	stock := producttestutils.GetMockRegistered(t, app)[0]
	prod := stock.GetInCostumerFormat()
	prod.Quantity = 40 // Set the desire quantity

	// Create the order OBJ
	ord, _ := costumerOrder.New(
		[]costumerOrder.CostumerProduct{*prod},
	)
	// Set the user that made the order request
	ord.UserID = usr.GetIdentifier()

	// Register the new order into DB
	_, err := app.Services.CostumerOrder.Register(ord)
	assert.NoError(t, err)
}

func TestRegister_InexistentUser(t *testing.T) {
	app := testutils.SetupTest(t)

	ord, _, _ := ordercostumertestutils.GetUnregisteredMock(t, app)

	_, err := app.Services.CostumerOrder.Register(ord)
	assert.ErrorIs(t, err, user_repository.ErrUserNotFound)
}

func TestRegister_InexistentProduct(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered user in DB
	usr := usertestutils.GetMockRegistered(t, app)

	// Get a registered product in DB
	stock := producttestutils.GetMock()[0]
	prod := stock.GetInCostumerFormat()
	prod.Quantity = 10 // Set the desire quantity

	// Create the order OBJ
	ord, _ := costumerOrder.New(
		[]costumerOrder.CostumerProduct{*prod},
	)
	// Set the user that made the order request
	ord.UserID = usr.GetIdentifier()

	// Register the new order into DB
	_, err := app.Services.CostumerOrder.Register(ord)
	assert.ErrorIs(t, err, product.ErrProductInexistent)
}

func TestRegister_InsufficientStock(t *testing.T) {
	app := testutils.SetupTest(t)

	// Get a registered user in DB
	usr := usertestutils.GetMockRegistered(t, app)

	// Get a registered product in DB
	stock := producttestutils.GetMockRegistered(t, app)[0]
	prod := stock.GetInCostumerFormat()
	prod.Quantity = 10000 // Set the desire quantity

	// Create the order OBJ
	ord, _ := costumerOrder.New(
		[]costumerOrder.CostumerProduct{*prod},
	)
	// Set the user that made the order request
	ord.UserID = usr.GetIdentifier()

	// Register the new order into DB
	_, err := app.Services.CostumerOrder.Register(ord)
	assert.ErrorIs(t, err, costumerOrder_service.ErrInsufficientStock)
}
