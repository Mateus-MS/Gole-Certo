package user_service

import (
	"context"

	user_cache "alves.com/modules/users/cache"
	user_model "alves.com/modules/users/model"
	user_repository "alves.com/modules/users/repo"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	Repo() *user_repository.Repository // TODO: REMOVE THIS
	Cache() *user_cache.Cache

	// Read
	ReadByName(context.Context, string) (user_model.UserEntity, error)

	Login(context.Context, string, string) (string, error)
	Register(context.Context, string, string) error
}

type service struct {
	repository *user_repository.Repository
	cache      *user_cache.Cache
}

func New(coll *mongo.Collection, cache *redis.Client, prefix string) *service {
	return &service{
		repository: user_repository.New(coll),
		cache:      user_cache.New(cache, prefix),
	}
}

func (s *service) Repo() *user_repository.Repository {
	return s.repository
}

func (s *service) Cache() *user_cache.Cache {
	return s.cache
}
