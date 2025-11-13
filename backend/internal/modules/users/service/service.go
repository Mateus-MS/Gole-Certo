package user_service

import (
	"context"
	"time"

	user_cache "alves.com/modules/users/cache"
	user_model "alves.com/modules/users/model"
	user_repository "alves.com/modules/users/repo"
	user_repository_mongo "alves.com/modules/users/repo/mongo"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	// Basic login
	Login(context.Context, string, string) (string, error)
	Register(context.Context, string, string) error

	// Cache functions
	SaveInCache(context.Context, string, user_model.UserCache, time.Duration) error
	ReadFromCache(context.Context, string) (user_model.UserCache, error)

	// DB crud functions to be exported
	Create(context.Context, user_model.UserEntity) error
	ReadByName(context.Context, string) (user_model.UserEntity, error)
	DeleteByID(context.Context, primitive.ObjectID) error
}

type service struct {
	repository user_repository.IRepository
	cache      *user_cache.Cache
}

func New(coll *mongo.Collection, cache *redis.Client, prefix string) *service {
	return &service{
		repository: user_repository_mongo.New(coll),
		cache:      user_cache.New(cache, prefix),
	}
}
