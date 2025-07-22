package costumerOrder_service_test

import (
	"testing"

	testutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test"
	ordercostumertestutils "github.com/Mateus-MS/Gole-Certo/dev/features/utils/test/orders/costumer"
	"github.com/stretchr/testify/assert"
)

func TestCreate_Success(t *testing.T) {
	app := testutils.SetupTest(t)

	ord, _, _ := ordercostumertestutils.GetUnregisteredMock(t, app)

	_, err := app.Services.CostumerOrder.Create(ord)
	assert.NoError(t, err)
}
