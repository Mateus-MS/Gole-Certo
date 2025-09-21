package main

import (
	stock_service "alves.com/backend/modules/stock/service"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := StartDBConnection()
	router := gin.Default()

	// SERVICES
	stockService := stock_service.New(db.Database("goleCertoDB").Collection("stock"))
	userService := user_service.New(db.Database("users").Collection("users"))

	app := NewApp(
		db,
		router,
		&Services{
			Stock: stockService,
			User:  userService,
		},
	)

	addRoutes(app)

	app.Router.Run("localhost:9090")
}
