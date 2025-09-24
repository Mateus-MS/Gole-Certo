package integration_users_test

import (
	"net/http"
	"testing"

	integration_helper "alves.com/tests/integration/helper"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin_Success(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	// Assuming that the register will work
	AttemptRegister(router, validUsername, validPassword)
	w := AttemptLogin(router, validUsername, validPassword)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
}

func TestUserLogin_Unregistered(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	w := AttemptLogin(router, validUsername, validPassword)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected HTTP 401")
}
