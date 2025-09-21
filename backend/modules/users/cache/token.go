package user_cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

// TODO: Seems strange and maybe a big security flaw
// This will map: token => duration
func (c *Cache) Set(ctx context.Context, token string, expiresAt time.Time, TTL time.Duration) error {
	return c.Redis.Set(ctx, token, expiresAt, TTL).Err()
}

func (c *Cache) Get(ctx context.Context, token string) (string, error) {
	val, err := c.Redis.Get(ctx, token).Result()
	if err == redis.Nil {
		return "", ErrTokenNotFound
	}
	return val, err
}
