package test_helper_users

import (
	"encoding/json"
	"time"

	generic_persistent "alves.com/backend/modules/common/model"
	user_model "alves.com/backend/modules/users/model"
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
