package userservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
)

type service struct {
	repository repository.UserRepository
}

func New(repo repository.UserRepository) *service {
	return &service{repository: repo}
}

func (s *service) Register(usr user.User) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = usr.IsValid(); err != nil {
		return err
	}

	return s.repository.Save(usr)
}

func (s *service) Search(identifier string) (usr user.User, err error) {
	return s.repository.Search(identifier)
}
