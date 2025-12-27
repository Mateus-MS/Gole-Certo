package stock_routes

import (
	"net/http"

	stock_model "alves.com/backend/modules/stock/model"
	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func StockCreate(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request stock_model.StockEntity

		// Parse the received json in the body
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := stockService.Create(c.Request.Context(), request)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		}
	}
}
