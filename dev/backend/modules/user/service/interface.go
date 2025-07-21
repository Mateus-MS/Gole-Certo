package user_service

import user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"

type Service interface {
	Create(user.User) error
	Read(string) (user.User, error)
}
