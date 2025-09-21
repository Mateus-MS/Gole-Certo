package stock_routes

import (
	"net/http"

	stock_model "alves.com/backend/modules/stock/model"
	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputStock stock_model.StockEntity

		err := c.ShouldBindJSON(&inputStock)

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err = stockService.Repo().Create(c, &inputStock)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "Product added to stock successfully")
	}
}
