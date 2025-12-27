package stock_routes

import (
	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func StockRead(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		// Use it however you want, e.g., query the stock service
		stock, err := stockService.ReadByName(c.Request.Context(), name)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}

		c.JSON(200, stock)
	}
}
