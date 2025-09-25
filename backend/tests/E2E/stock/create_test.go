package e2e_stock_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	user_service "alves.com/modules/users/service"
	test_helper_app "alves.com/tests/helper"
	test_helper_stock "alves.com/tests/helper/services/stock"
	"github.com/stretchr/testify/assert"
)

func TestStockCreate_WithoutBearerHeader(t *testing.T) {
	t.Parallel()
	app := test_helper_app.NewApp(t)

	// Get the product
	productJson := test_helper_stock.GetProductJson("Coca cola")

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(productJson))
	req.Header.Set("Content-Type", "application/json")

	// Sent the request
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
	token, _ := user_service.GenerateRandomToken(20)
	req.Header.Set("Authorization", "Bearer "+token)

	// Sent the request
	w := httptest.NewRecorder()
	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected HTTP 401")
}
