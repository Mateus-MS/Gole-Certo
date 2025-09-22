package user_routes

import (
	"net/http"

	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the credentials
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Try to login the user with the received credentials
		token, err := userService.Login(c, username, password)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Return token in JSON instead of cookie
		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
			"message":     "logged in successfully",
		})
	}
}
