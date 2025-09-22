package user_routes

import (
	"errors"
	"net/http"

	generic_repository "alves.com/modules/common/repo"
	user_repository "alves.com/modules/users/repo"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
)

func UserDelete(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "user id not found")
			return
		}
		err := userService.Repo().DeleteByID(c, userID.(string))
		if err != nil {
			if errors.Is(err, user_repository.ErrUserInexistent) || errors.Is(err, generic_repository.ErrItemInexistent) {
				c.String(http.StatusNotFound, "this user doesn't exists")
				return
			}

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "user deleted successfully")
	}
}
