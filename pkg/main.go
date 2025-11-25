package main

import (
	"log"

	"alves.com/backend/app"
	"alves.com/backend/app/config"
	"alves.com/backend/app/routes"
	stock_service "alves.com/backend/modules/stock/service"
	user_service "alves.com/backend/modules/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	db := config.StartDBConnection()
	cache := config.StartCacheConnection()
	router := gin.Default()

	// SERVICES
	stockService := stock_service.New(db.Database("cluster").Collection("stock"))
	userService := user_service.New(db.Database("cluster").Collection("users"), cache, "")

	aplication := app.NewApp(
		db,
		cache,
		router,
		&app.Services{
			Stock: stockService,
			User:  userService,
		},
	)

	routes.InitRoutes(aplication)

	aplication.Router.Run(":8080")
}
