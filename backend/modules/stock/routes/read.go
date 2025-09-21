package stock_routes

import (
	"errors"
	"fmt"
	"net/http"

	stock_error "alves.com/backend/modules/stock/errors"
	stock_service "alves.com/backend/modules/stock/service"
	"github.com/gin-gonic/gin"
)

func ReadProduct(stockService stock_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		product, err := stockService.Repo().ReadByName(c, name)

		if err != nil {
			if errors.Is(err, stock_error.ErrProductInexistent) {
				c.String(404, err.Error())
				return
			}
			if errors.Is(err, stock_error.ErrCannotConvert) {
				c.String(500, err.Error())
				return
			}

			c.String(500, fmt.Errorf("Something went wrong: %w", err).Error())
		}

		c.JSON(http.StatusOK, product)
	}
}
