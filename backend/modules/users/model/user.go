package user_model

import generic_persistent "alves.com/backend/modules/common/model"

type UserEntity struct {
	generic_persistent.Persistent `bson:",inline"`

	Name         string `json:"name"         binding:"required"`
	Password     string `json:"password"     binding:"required"`
	SessionToken string `json:"sessionToken" binding:"required"`
}

func (u *UserEntity) GetDTO() *UserDTO {
	return &UserDTO{
		Name: u.Name,
	}
}

type UserDTO struct {
	Name string `json:"name"`
}
