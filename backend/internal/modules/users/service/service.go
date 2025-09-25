package user_service

import (
	"context"
	"time"

	user_cache "alves.com/modules/users/cache"
	user_model "alves.com/modules/users/model"
	user_repository "alves.com/modules/users/repo"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: if it became to get clunky, split this interface in parts, for a better organization
type IService interface {
	// Register(context.Context, product.ProductStock) error

	// DeductFromStock(context.Context, product.ProductStock, int64) error

	// DB crud functions to be exported
	Create(context.Context, user_model.UserEntity) error

	ReadByName(context.Context, string) (user_model.UserEntity, error)

	UpdateByName(context.Context, user_model.UserEntity) error

	DeleteByID(context.Context, primitive.ObjectID) error

	// Basic login
	Login(context.Context, string, string) (string, error)
	Register(context.Context, string, string) error

	// Cache functions
	SaveInCache(context.Context, string, user_model.UserCache, time.Duration) error
	ReadFromCache(context.Context, string) (user_model.UserCache, error)
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
