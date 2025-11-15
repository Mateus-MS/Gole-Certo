package routes

import (
	"alves.com/backend/app"
	"alves.com/backend/middlewares"
	stock_routes "alves.com/backend/modules/stock/routes"
	user_routes "alves.com/backend/modules/users/routes"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *app.App) {
	app.Router.POST("/products", middlewares.AuthMiddleware(app.Services.User), middlewares.IsAdmin(), stock_routes.CreateProduct(app.Services.Stock))
	app.Router.GET("/products", stock_routes.ReadProduct(app.Services.Stock))

	RegisterUserRoutes(app.Router, app.Services.User)
}

func RegisterUserRoutes(router *gin.Engine, serv user_service.IService) {
	users := router.Group("/users")

	users.GET("/:name", user_routes.UserRead(serv))
	users.GET("/protected", middlewares.AuthMiddleware(serv), user_routes.UserProtected(serv))

	users.POST("/login", user_routes.UserLogin(serv))
	users.POST("/register", user_routes.UserRegister(serv))
	users.POST("/delete", middlewares.AuthMiddleware(serv), user_routes.UserDelete(serv))
}
