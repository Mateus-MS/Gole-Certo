package test_helper_app

import (
	"log"
	"testing"

	"alves.com/backend/app"
	"alves.com/backend/app/routes"
	stock_service "alves.com/backend/modules/stock/service"
	user_cache "alves.com/backend/modules/users/cache"
	user_service "alves.com/backend/modules/users/service"
	tests_mock "alves.com/backend/tests/helper/mock"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestDependencies struct {
	Router *gin.Engine
	DB     *mongo.Database
	Cache  *user_cache.Cache
}

func NewAppBase(t *testing.T) *TestDependencies {
	t.Helper()
	gin.SetMode(gin.TestMode)
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Println("Warning: .env file not loaded")
	}

	database := tests_mock.SetupDB(t)
	cache := tests_mock.SetupCache(t)

	router := gin.Default()

	return &TestDependencies{Router: router, DB: database, Cache: cache}
}

func NewApp(t *testing.T) *app.App {
	dependencies := NewAppBase(t)

	services := app.Services{
		Stock: stock_service.New(dependencies.DB.Collection("stock")),
		User:  user_service.New(dependencies.DB.Collection("users"), dependencies.Cache.Redis, dependencies.Cache.Prefix),
	}

	application := app.NewApp(dependencies.DB.Client(), dependencies.Cache.Redis, dependencies.Router, &services)

	routes.InitRoutes(application)

	return application
}
