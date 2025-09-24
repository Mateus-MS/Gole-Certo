package integration_users_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var validUsername = "jhonDoe"
var validPassword = "jhonpass"

func AttemptLogin(router *gin.Engine, username, password string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/users/login", nil)
	req.SetBasicAuth(username, password)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func AttemptRegister(router *gin.Engine, username, password string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/users/register", nil)
	req.SetBasicAuth(username, password)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func AttemptRead(router *gin.Engine, username string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/users/"+username, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
