package user_service

import (
	"context"

	user_repository "alves.com/backend/modules/users/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	Repo() *user_repository.Repository

	Login(context.Context, string, string) error
}

type service struct {
	repository *user_repository.Repository
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: user_repository.New(coll),
	}
}

func (s *service) Repo() *user_repository.Repository {
	return s.repository
}
