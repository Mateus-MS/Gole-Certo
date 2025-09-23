package integration_users

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

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
