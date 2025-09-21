package main

import (
	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := StartDBConnection()
	router := gin.Default()

	// SERVICES
	stockService := stock_service.New(db.Database("goleCertoDB").Collection("stock"))

	app := NewApp(
		db,
		router,
		&Services{
			Stock: stockService,
		},
	)

	addRoutes(app)

	app.Router.Run("localhost:9090")
}
