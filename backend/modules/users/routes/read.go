package user_routes

import (
	"errors"
	"fmt"
	"net/http"

	user_error "alves.com/backend/modules/users/errors"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func ReadUser(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		user, err := userService.Repo().ReadByName(c, name)

		if err != nil {
			if errors.Is(err, user_error.ErrUserInexistent) {
				c.String(404, err.Error())
				return
			}
			if errors.Is(err, user_error.ErrCannotConvert) {
				c.String(500, err.Error())
				return
			}

			c.String(500, fmt.Errorf("Something went wrong: %w", err).Error())
		}

		c.JSON(http.StatusOK, user)
	}
}
