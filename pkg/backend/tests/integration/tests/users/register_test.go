package integration_users_test

import (
	"testing"

	integration_setup "alves.com/backend/tests/integration/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRegister_Success(t *testing.T) {
	t.Parallel()
	h := integration_setup.NewHarness(t)

	var validUsername = "jhonDoe"
	var validPassword = "jhonPass"

	// Try register
	{
		err := h.Services.User.Register(h.Ctx, validUsername, validPassword)
		require.NoError(t, err, "user register should not return an error")
	}

	// Try query by the name
	{
		userEntity, err := h.Services.User.ReadByName(h.Ctx, validUsername)
		assert.NoError(t, err, "user query should not return an error")
		assert.Equal(t, validUsername, userEntity.Name, "username query should be the same as registered")
	}
}
