package user_service

import (
	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	user_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository *user_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{repository: user_repository.New(coll)}
}

func (s *service) Register(usr user.User) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = usr.IsValid(); err != nil {
		return err
	}

	return s.repository.Create(usr)
}

func (s *service) Repo() *user_repository.Repository {
	return s.repository
}
