package userservice

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"

type Service interface {
	Register(user.User) error
	Search(string) (user.User, error)
}
