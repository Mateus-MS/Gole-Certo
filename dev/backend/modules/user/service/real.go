package user_service

import (
	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	user_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository user_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{repository: *user_repository.New(coll)}
}

// TODO: instead of usr having a field `type`, it "discovers" here
func (s *service) Create(usr user.User) (err error) {
	// TODO: See if is need to first check if already exists a client equals to the received one
	if err = usr.IsValid(); err != nil {
		return err
	}

	return s.repository.Create(usr)
}

func (s *service) Read(identifier string) (usr user.User, err error) {
	return s.repository.Read(identifier)
}
