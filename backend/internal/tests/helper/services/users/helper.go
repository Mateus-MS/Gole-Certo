package test_helper_users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	generic_persistent "alves.com/modules/common/model"
	user_model "alves.com/modules/users/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserJson(name string) []byte {
	prodEntity := user_model.UserEntity{
		Persistent: generic_persistent.Persistent{
			ID: primitive.NewObjectIDFromTimestamp(time.Now()),
		},

		Name: name,
	}
	jsonData, _ := json.Marshal(prodEntity)

	return jsonData
}

func AttemptRegister(router *gin.Engine, username, password string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/users/register", nil)
	req.SetBasicAuth(username, password)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func AttemptLogin(router *gin.Engine, username, password string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodPost, "/users/login", nil)
	req.SetBasicAuth(username, password)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
