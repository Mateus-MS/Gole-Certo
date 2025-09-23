package integration_users_test

import (
	"net/http"
	"testing"

	integration_helper "alves.com/tests/integration/helper"
	integration_users "alves.com/tests/integration/users"
	"github.com/stretchr/testify/assert"
)

var validUsername = "jhonDoe"
var validPassword = "jhonpass"

func TestUserGET_Success(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	integration_users.AttemptRegister(router, validUsername, validPassword)

	w := integration_users.AttemptRead(router, validUsername)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")

	println(w.Body.String())
}
