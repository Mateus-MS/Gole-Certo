package integration_users_test

import (
	"net/http"
	"testing"

	integration_helper "alves.com/tests/integration/helper"
	"github.com/stretchr/testify/assert"
)

func TestUserGET_Success(t *testing.T) {
	t.Parallel()
	router := integration_helper.SetupUserApp(t)

	AttemptRegister(router, validUsername, validPassword)

	w := AttemptRead(router, validUsername)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")

	println(w.Body.String())
}
