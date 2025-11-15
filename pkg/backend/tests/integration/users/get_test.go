package integration_users_test

import (
	"testing"

	user_service "alves.com/backend/modules/users/service"
	test_helper_app "alves.com/backend/tests/helper"
	"github.com/stretchr/testify/assert"
)

var validUsername = "jhonDoe"
var validPassword = "jhonPass"

func TestUserGET_Success(t *testing.T) {
	t.Parallel()
	deps := test_helper_app.NewAppBase(t)
	userService := user_service.New(deps.DB.Collection("users"), deps.Cache.Redis, deps.Cache.Prefix)

	// Try to register
	{
		err := userService.Register(t.Context(), validUsername, validPassword)
		assert.Nil(t, err, "user registration should not return an error")
	}

	// Try to read the user
	{
		_, err := userService.ReadByName(t.Context(), validUsername)
		assert.Nil(t, err, "user query should not return an error")
	}
}
