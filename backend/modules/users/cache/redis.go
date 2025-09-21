package user_cache

import (
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Redis *redis.Client
}

func New() *Cache {
	return &Cache{
		Redis: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
