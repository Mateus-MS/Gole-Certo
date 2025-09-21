package stock_routes

import (
	"net/http"

	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			},
		)
	}
}
