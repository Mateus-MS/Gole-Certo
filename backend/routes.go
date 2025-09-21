package main

import (
	"alves.com/backend/middlewares"
	stock_routes "alves.com/backend/modules/stock/routes"
	user_routes "alves.com/backend/modules/users/routes"
)

func addRoutes(app *App) {
	admCollection := app.DB.Database("users").Collection("adms")

	app.Router.POST("/products", middlewares.IsAdmin(admCollection), stock_routes.CreateProduct(app.Services.Stock))
	app.Router.GET("/products", stock_routes.ReadProduct(app.Services.Stock))

	app.Router.GET("/users/:name", user_routes.UserRead(app.Services.User))
	app.Router.POST("/users/login", user_routes.UserLogin(app.Services.User))
}
