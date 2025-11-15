package e2e_users_test

import (
	"net/http"
	"testing"

	test_helper_app "alves.com/backend/tests/helper"
	test_helper_users "alves.com/backend/tests/helper/services/users"
	"github.com/stretchr/testify/assert"
)

var validUsername = "jhonDoe"
var validPassword = "jhonPass"

func TestUserLogin_Success(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Register an user
	{
		w := test_helper_users.AttemptRegister(app.Router, validUsername, validPassword)
		assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
	}

	// Try to log in
	{
		w := test_helper_users.AttemptLogin(app.Router, validUsername, validPassword)
		assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
	}

}
