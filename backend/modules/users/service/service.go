package user_service

import (
	"context"

	user_cache "alves.com/modules/users/cache"
	user_repository "alves.com/modules/users/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	Repo() *user_repository.Repository
	Cache() *user_cache.Cache

	Login(context.Context, string, string) (string, error)
	Register(context.Context, string, string) error
}

type service struct {
	repository *user_repository.Repository
	cache      *user_cache.Cache
}

func New(coll *mongo.Collection) *service {
	return &service{
		repository: user_repository.New(coll),
		cache:      user_cache.New(),
	}
}

func (s *service) Repo() *user_repository.Repository {
	return s.repository
}

func (s *service) Cache() *user_cache.Cache {
	return s.cache
}
