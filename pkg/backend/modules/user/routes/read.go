package user_routes

import (
	"errors"
	"fmt"
	"net/http"

	user_error "alves.com/backend/modules/user/errors"
	user_service "alves.com/backend/modules/user/service"
	"github.com/gin-gonic/gin"
)

func UserRead(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		user, err := userService.ReadByName(c, name)

		if err != nil {
			if errors.Is(err, user_error.ErrUserInexistent) {
				c.String(404, err.Error())
				return
			}

			c.String(500, fmt.Errorf("something went wrong: %w", err).Error())
		}

		c.JSON(http.StatusOK, user.GetDTO())
	}
}
