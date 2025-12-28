package integration_users_test

import (
	"testing"

	user_error "alves.com/backend/modules/user/errors"
	integration_setup "alves.com/backend/tests/integration/setup"
	integration_fixtures "alves.com/backend/tests/integration/setup/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestUserRead_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	var validUsername = "jhonDoe"
	var validPassword = "jhonPass"

	// Direct insert the user into DB
	integration_fixtures.InsertTestUser(h.Ctx, h.DB.Database, validUsername, validPassword)

	// Try to read the user
	{
		_, err := h.Services.User.ReadByName(h.Ctx, validUsername)
		assert.NoError(t, err, "user query should not return an error")
	}
}

func TestUserRead_Inexistent(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	var validUsername = "jhonDoe"

	// Try to read a inexistent user
	{
		_, err := h.Services.User.ReadByName(h.Ctx, validUsername)
		assert.ErrorIs(t, err, user_error.ErrUserInexistent, "user query should return a not found error")
	}
}
