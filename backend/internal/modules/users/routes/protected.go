package user_routes

import (
	"net/http"

	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
)

func UserProtected(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("userID")
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.JSON(http.StatusOK, "Successfully accessed the protected route")
	}
}
