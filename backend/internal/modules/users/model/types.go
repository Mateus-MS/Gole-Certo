package user_model

import (
	generic_persistent "alves.com/modules/common/model"
)

// Represents the way that the user is stored at DB layer
type UserEntity struct {
	generic_persistent.Persistent `bson:",inline"`

	Name     string `json:"name"         binding:"required" bson:"name"`
	Password string `json:"password"     binding:"required" bson:"password"`
	IsAdmin  bool   `json:"isAdmin"                         bson:"isAdmin"`
}

// Represents the way that the user is stored in Cache layer
type UserCache struct {
	generic_persistent.Persistent
	IsAdmin bool
}

// Represents the way that the user is sented to clients, ommiting senstive data
type UserDTO struct {
	Name string `json:"name"`
}
