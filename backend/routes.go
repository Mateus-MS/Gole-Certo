package main

import (
	stock_routes "alves.com/backend/modules/stock/routes"
)

func addRoutes(app *App) {
	app.Router.GET("/product", stock_routes.CreateProduct(app.Services.Stock))
}
