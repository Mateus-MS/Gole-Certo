package e2e_stock_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"alves.com/pkg/security"
	test_helper_app "alves.com/tests/helper"
	test_helper_stock "alves.com/tests/helper/services/stock"
	test_helper_users "alves.com/tests/helper/services/users"
	"github.com/stretchr/testify/assert"
)

func TestStockCreate_Success(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Get the ADM accessToken
	accessToken := test_helper_users.LoginTempADM(t, app.Services.User)

	// Get the product
	productJson := test_helper_stock.GetProductJson("Coca cola")

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJson))
	req.Header.Set("Content-Type", "application/json")

	// Add the adm accessToken to header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send create new product request
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")

	// Try to read the just created product
	req, _ = http.NewRequest(http.MethodGet, "/products?name=Coca cola", nil)
	w = httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected HTTP 200")
}

func TestStockCreate_WithoutBearerHeader(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Get the product
	productJson := test_helper_stock.GetProductJson("Coca cola")

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJson))
	req.Header.Set("Content-Type", "application/json")

	// Send request
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "expected HTTP 400")
}

func TestStockCreate_WithInvalidBearerToken(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Get the product
	productJson := test_helper_stock.GetProductJson("Coca cola")

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJson))
	req.Header.Set("Content-Type", "application/json")

	// Add an invalid token to header
	token, _ := security.GenerateRandomToken(20)
	req.Header.Set("Authorization", "Bearer "+token)

	// Send request
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected HTTP 401")
}
