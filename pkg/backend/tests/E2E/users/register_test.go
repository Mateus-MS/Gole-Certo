package e2e_users_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_helper_app "alves.com/backend/tests/helper"
	test_helper_users "alves.com/backend/tests/helper/services/users"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister_Success(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/users/register", nil)
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(validUsername, validPassword)

	// Sent the request
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
}

func TestUserRegister_DuplicatedUser(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// First request
	{
		w := test_helper_users.AttemptRegister(app.Router, validUsername, validPassword)
		assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
	}

	// Duplicated request
	{
		w := test_helper_users.AttemptRegister(app.Router, validUsername, validPassword)
		assert.Equal(t, http.StatusConflict, w.Code, "expected HTTP 409")
	}
}
