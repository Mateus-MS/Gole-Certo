package routes

import (
	"alves.com/backend/app"
	"alves.com/backend/middlewares"
	order_routes "alves.com/backend/modules/order/routes"
	order_service "alves.com/backend/modules/order/service"
	stock_routes "alves.com/backend/modules/stock/routes"
	user_routes "alves.com/backend/modules/user/routes"
	user_service "alves.com/backend/modules/user/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *app.App) {
	app.Router.POST("/products", middlewares.AuthMiddleware(app.Services.User), middlewares.IsAdmin(), stock_routes.StockCreate(app.Services.Stock))
	app.Router.GET("/products", stock_routes.StockRead(app.Services.Stock))

	RegisterUserRoutes(app.Router, app.Services.User)
	RegisterOrderRoutes(app.Router, app.Services.Order)
}

func RegisterUserRoutes(router *gin.Engine, serv user_service.IService) {
	users := router.Group("/users")

	users.GET("/:name", user_routes.UserRead(serv))
	users.GET("/protected", middlewares.AuthMiddleware(serv), user_routes.UserProtected(serv))

	users.POST("/login", user_routes.UserLogin(serv))
	users.POST("/register", user_routes.UserRegister(serv))
	users.POST("/delete", middlewares.AuthMiddleware(serv), user_routes.UserDelete(serv))
}

func RegisterOrderRoutes(router *gin.Engine, serv order_service.IService) {
	order := router.Group("/order")

	order.POST("/create", order_routes.OrderCreate(serv))
}
