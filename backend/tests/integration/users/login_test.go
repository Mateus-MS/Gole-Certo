package integration_users_test

import (
	"net/http"
	"testing"

	integration_helper "alves.com/tests/integration/helper"
	integration_users "alves.com/tests/integration/users"
	"github.com/stretchr/testify/assert"
)

var ValidUsername = "jhonDoe"
var ValidPassword = "jhonpass"

func TestUserLogin_Success(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	// Assuming that the register will work
	integration_users.AttemptRegister(router, ValidUsername, ValidPassword)
	w := integration_users.AttemptLogin(router, ValidUsername, ValidPassword)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
}

func TestUserLogin_Unregistered(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	w := integration_users.AttemptLogin(router, ValidUsername, ValidPassword)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected HTTP 401")
}
