package user_cache

import (
	"context"
	"errors"
	"strconv"
	"time"

	user_model "alves.com/backend/modules/user/model"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrTokenNotFound = errors.New("the given token is invalid or expired")
)

type Cache struct {
	Redis  *redis.Client
	Prefix string
}

func New(client *redis.Client, prefix string) *Cache {
	return &Cache{
		Redis:  client,
		Prefix: prefix,
	}
}

func (c *Cache) Save(ctx context.Context, token string, userData user_model.UserCache, TTL time.Duration) error {
	// Convert user data to a map for Redis
	data := map[string]interface{}{
		"userID":  userData.ID.Hex(),
		"isAdmin": userData.IsAdmin,
	}

	key := c.Prefix + token

	// Save the data into redis
	err := c.Redis.HSet(ctx, key, data).Err()
	if err != nil {
		return err
	}

	// Add an expiration date to it
	err = c.Redis.Expire(ctx, key, TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) Read(ctx context.Context, token string) (user_model.UserCache, error) {
	userData := user_model.UserCache{}

	key := c.Prefix + token

	data, err := c.Redis.HGetAll(ctx, key).Result()
	if err != nil {
		return user_model.UserCache{}, err
	}

	userData.ID, err = primitive.ObjectIDFromHex(data["userID"])
	if err != nil {
		return user_model.UserCache{}, err
	}

	userData.IsAdmin, err = strconv.ParseBool(data["isAdmin"])
	if err != nil {
		return user_model.UserCache{}, err
	}

	return userData, nil
}
