package user_cache

import (
	"alves.com/app/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Redis *redis.Client
}

func New() *Cache {
	println("redis URI: " + config.GetRedisURI())
	return &Cache{
		Redis: redis.NewClient(&redis.Options{
			Addr:     config.GetRedisURI(),
			Password: "",
			DB:       0,
		}),
	}
}
