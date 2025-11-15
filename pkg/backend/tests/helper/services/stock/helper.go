package test_helper_stock

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	generic_persistent "alves.com/backend/modules/common/model"
	stock_model "alves.com/backend/modules/stock/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validUsername = "jhonDoe"
var validPassword = "jhonpass"

func GetProductJson(name string) []byte {
	prodEntity := stock_model.StockEntity{
		Persistent: generic_persistent.Persistent{
			ID: primitive.NewObjectIDFromTimestamp(time.Now()),
		},

		Name: name,
	}
	jsonData, _ := json.Marshal(prodEntity)

	return jsonData
}

func AttemptCreate(router *gin.Engine, name string) *httptest.ResponseRecorder {
	// Create the data to be sent
	jsonData := GetProductJson(name)

	// Create the request
	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Sent the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
