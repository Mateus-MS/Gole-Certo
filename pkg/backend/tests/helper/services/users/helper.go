package test_helper_users

import (
	"encoding/json"
	"time"

	user_model "alves.com/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserJson(name string) []byte {
	prodEntity := user_model.UserEntity{
		ID:   primitive.NewObjectIDFromTimestamp(time.Now()),
		Name: name,
	}
	jsonData, _ := json.Marshal(prodEntity)

	return jsonData
}
