package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartDBConnection() (mongoClient *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	if mongoClient, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI(GetMongoURI())); err != nil {
		log.Fatal("Mongo connection error: " + err.Error())
	}

	return mongoClient
}

func StartCacheConnection() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     GetRedisURI(),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := redis.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis ping failed:", err)
	}

	return redis
}
