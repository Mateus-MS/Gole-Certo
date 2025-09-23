package integration_helper

import (
	"testing"

	"alves.com/app/routes"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
)

func SetupUserApp(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	database := SetupDB(t)
	userService := user_service.New(database.Collection("user"))

	router := gin.Default()
	routes.RegisterUserRoutes(router, userService)

	return router
}
