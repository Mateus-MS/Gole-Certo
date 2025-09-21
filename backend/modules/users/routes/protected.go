package user_routes

import (
	"net/http"

	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func UserProtected(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEntity, _ := c.Get("user")
		c.JSON(http.StatusOK, userEntity)
	}
}
