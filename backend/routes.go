package main

import (
	stock_routes "alves.com/backend/modules/stock/routes"
)

func addRoutes(app *App) {
	app.Router.POST("/products", stock_routes.CreateProduct(app.Services.Stock))
	app.Router.GET("/products/:name", stock_routes.ReadProduct(app.Services.Stock))
}
