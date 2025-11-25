package user_routes

import (
	"errors"
	"net/http"

	user_error "alves.com/backend/modules/user/errors"
	user_service "alves.com/backend/modules/user/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserDelete(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "user id not found")
			return
		}
		idObj, _ := primitive.ObjectIDFromHex(userID.(string))
		err := userService.DeleteByID(c, idObj)
		if err != nil {
			if errors.Is(err, user_error.ErrUserInexistent) {
				c.String(http.StatusNotFound, "this user doesn't exists")
				return
			}

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, "user deleted successfully")
	}
}
