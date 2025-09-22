package user_model

import (
	"encoding/json"
	"fmt"
	"time"

	generic_persistent "alves.com/modules/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	generic_persistent.Persistent `bson:",inline"`

	Name     string `json:"name"         binding:"required" bson:"name"`
	Password string `json:"password"     binding:"required" bson:"password"`
}

func NewUser(username, password string) *UserEntity {
	user := UserEntity{
		Name:     username,
		Password: password,
	}

	user.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	return &user
}

func (u *UserEntity) ToString() string {
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return fmt.Sprintf("error converting UserEntity to string: %v", err)
	}
	return string(data)
}

func (u *UserEntity) GetDTO() *UserDTO {
	return &UserDTO{
		Name: u.Name,
	}
}

type UserDTO struct {
	Name string `json:"name"`
}
