package stock_routes

import (
	"net/http"
	"time"

	stock_model "alves.com/modules/stock/model"
	stock_service "alves.com/modules/stock/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputStock stock_model.StockEntity

		err := c.ShouldBindJSON(&inputStock)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		inputStock.ID = primitive.NewObjectIDFromTimestamp(time.Now())

		err = stockService.Register(c, inputStock)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "Product added to stock successfully")
	}
}
