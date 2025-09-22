package main

import (
	stock_service "alves.com/modules/stock/service"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Stock stock_service.IService
	User  user_service.IService
}

type App struct {
	DB       *mongo.Client
	Router   *gin.Engine
	Services *Services
}

func NewApp(db *mongo.Client, router *gin.Engine, services *Services) *App {
	return &App{
		DB:       db,
		Router:   router,
		Services: services,
	}
}
