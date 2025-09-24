package app

import (
	stock_service "alves.com/modules/stock/service"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Stock stock_service.IService
	User  user_service.IService
}

type App struct {
	DB       *mongo.Client
	Cache    *redis.Client
	Router   *gin.Engine
	Services *Services
}

func NewApp(db *mongo.Client, cache *redis.Client, router *gin.Engine, services *Services) *App {
	return &App{
		DB:       db,
		Cache:    cache,
		Router:   router,
		Services: services,
	}
}
