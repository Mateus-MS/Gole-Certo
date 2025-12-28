package user_model

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(username, hashedPassword string) *UserEntity {
	user := UserEntity{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:     username,
		Password: hashedPassword,
		IsAdmin:  false,
	}

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
		ID:      u.ID,
		IsAdmin: u.IsAdmin,
	}
}
