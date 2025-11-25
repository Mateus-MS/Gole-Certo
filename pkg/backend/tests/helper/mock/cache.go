package tests_mock

import (
	"context"
	"fmt"
	"testing"
	"time"

	"alves.com/backend/app/config"
	user_cache "alves.com/backend/modules/user/cache"
	"github.com/redis/go-redis/v9"
)

func SetupCache(t *testing.T) *user_cache.Cache {
	t.Helper()

	prefix := fmt.Sprintf("test_%s:", t.Name())
	cache := &user_cache.Cache{
		Redis: redis.NewClient(&redis.Options{
			Addr: config.GetRedisURI(),
			DB:   0,
		}),
		Prefix: prefix,
	}

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Remove all keys with this prefix
		keys, _ := cache.Redis.Keys(ctx, prefix+"*").Result()
		if len(keys) > 0 {
			_ = cache.Redis.Del(ctx, keys...).Err()
		}

		_ = cache.Redis.Close()
	})

	return cache
}
