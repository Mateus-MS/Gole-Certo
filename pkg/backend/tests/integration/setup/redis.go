package integration_setup

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func NewTestRedis(t *testing.T) *redis.Client {
	t.Helper()

	// Assign a unique DB index per test based on test name hash
	dbIndex := hashDBIndex(t.Name())

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // your running Redis
		DB:   dbIndex,
	})

	// Flush this DB to ensure a clean state
	if err := client.FlushDB(context.Background()).Err(); err != nil {
		t.Fatalf("failed to flush Redis: %v", err)
	}

	return client
}

// simple hash to pick DB index 0..15 (adjust as needed)
func hashDBIndex(name string) int {
	var sum int
	for _, c := range name {
		sum += int(c)
	}
	return sum % 16
}
