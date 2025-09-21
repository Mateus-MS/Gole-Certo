package user_routes

import (
	"errors"
	"fmt"
	"net/http"

	user_repository "alves.com/backend/modules/users/repo"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

func UserRead(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		user, err := userService.Repo().ReadByName(c, name)

		if err != nil {
			if errors.Is(err, user_repository.ErrUserInexistent) {
				c.String(404, err.Error())
				return
			}
			if errors.Is(err, user_repository.ErrCannotConvert) {
				c.String(500, err.Error())
				return
			}

			c.String(500, fmt.Errorf("Something went wrong: %w", err).Error())
		}

		c.JSON(http.StatusOK, user.GetDTO())
	}
}
