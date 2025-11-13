package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("userIsAdmin")
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
