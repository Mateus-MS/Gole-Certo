package user_routes

import (
	"errors"
	"net/http"

	user_error "alves.com/backend/modules/user/errors"
	user_service "alves.com/backend/modules/user/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := userService.Register(c, username, password)
		if err != nil {
			if errors.Is(err, user_error.ErrUserAlreadyExists) {
				c.String(http.StatusConflict, "username already registered")
				return
			}

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "User registered into DB successfully")
	}
}
