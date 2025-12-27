package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, ok := c.Get("userIsAdmin")
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// type assert and check value
		if adminBool, ok := isAdmin.(bool); !ok || !adminBool {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
