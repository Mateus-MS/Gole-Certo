package user_service

import (
	user_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	repository *user_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{repository: user_repository.New(coll)}
}

func (s *service) Repo() *user_repository.Repository {
	return s.repository
}
