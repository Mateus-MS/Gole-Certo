package userservice

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"

type Service interface {
	Create(user.User) error
	Read(string) (user.User, error)
}
