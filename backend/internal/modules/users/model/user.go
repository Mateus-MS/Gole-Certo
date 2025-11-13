package user_model

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (u *UserEntity) GetCache() *UserCache {
	return &UserCache{
		Persistent: u.Persistent,
		IsAdmin:    u.IsAdmin,
	}
}
