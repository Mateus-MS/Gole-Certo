package config

import "os"

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetRedisURI() string {
	return os.Getenv("REDIS_ADDR")
}
