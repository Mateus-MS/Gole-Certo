package integration_helper

import (
	"log"
	"testing"

	"alves.com/app/routes"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupUserApp(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Println("Warning: .env file not loaded")
	}

	database := SetupDB(t)
	cache := SetupCache(t)
	// Kinda dumb but will stay like this till it bothers me enough
	userService := user_service.New(database.Collection("user"), cache.Redis, cache.Prefix)

	router := gin.Default()
	routes.RegisterUserRoutes(router, userService)

	return router
}
