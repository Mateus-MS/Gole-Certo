package integration_helper

import (
	"context"
	"fmt"
	"testing"
	"time"

	"alves.com/app/config"
	user_cache "alves.com/modules/users/cache"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(t *testing.T) *mongo.Database {
	t.Helper()

	// Create an unique name to this DB
	dbName := fmt.Sprintf("testdb_%s", t.Name())

	client, err := mongo.Connect(t.Context(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database(dbName)

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = client.Database(dbName).Drop(ctx)
		_ = client.Disconnect(ctx)
	})

	return db
}

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
