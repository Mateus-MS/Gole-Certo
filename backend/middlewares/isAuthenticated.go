package middlewares

import (
	"errors"
	"net/http"
	"strings"

	user_cache "alves.com/modules/users/cache"
	user_service "alves.com/modules/users/service"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissingAuthHeader = errors.New("missing authorization header")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
	ErrExpiredToken      = errors.New("expired token provided")
)

func AuthMiddleware(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the token from the header
		token, err := getTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, ErrMissingAuthHeader) || errors.Is(err, ErrInvalidAuthHeader) {
				status = http.StatusBadRequest
			}
			c.String(status, err.Error())
			return
		}

		// Check if the token is on Cache
		userID, err := userService.Cache().Get(c, token)
		if err != nil {
			if errors.Is(err, user_cache.ErrTokenNotFound) {
				c.String(http.StatusUnauthorized, "not logged")
				return
			}

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func getTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", ErrMissingAuthHeader
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}
