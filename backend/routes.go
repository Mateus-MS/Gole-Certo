package main

import (
	"alves.com/middlewares"
	stock_routes "alves.com/modules/stock/routes"
	user_routes "alves.com/modules/users/routes"
)

func addRoutes(app *App) {
	admCollection := app.DB.Database("users").Collection("adms")

	app.Router.POST("/products", middlewares.IsAdmin(admCollection), stock_routes.CreateProduct(app.Services.Stock))
	app.Router.GET("/products", stock_routes.ReadProduct(app.Services.Stock))

	app.Router.GET("/users/:name", user_routes.UserRead(app.Services.User))
	app.Router.POST("/users/login", user_routes.UserLogin(app.Services.User))
	app.Router.POST("/users/register", user_routes.UserRegister(app.Services.User))
	app.Router.GET("/users/protected", middlewares.AuthMiddleware(app.Services.User), user_routes.UserProtected(app.Services.User))
	app.Router.POST("/users/delete", middlewares.AuthMiddleware(app.Services.User), user_routes.UserDelete(app.Services.User))
}
