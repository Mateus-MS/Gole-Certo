package middlewares

import (
	"errors"
	"net/http"
	"strings"

	user_repository "alves.com/backend/modules/users/repo"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissingAuthHeader = errors.New("missing authorization header")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
	ErrExpiredToken      = errors.New("expired token provided")
)

func AuthMiddleware(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		token, err := getTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			if errors.Is(err, ErrMissingAuthHeader) {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Query for a user with that token
		userEntity, err := userService.Repo().ReadBySessionToken(c, token)
		if err != nil {
			if errors.Is(err, user_repository.ErrUserInexistent) {
				c.String(http.StatusUnauthorized, "none user match the provided credential")
				return
			}

			c.String(http.StatusInternalServerError, "some ")
			return
		}

		// Validate the token
		if !userEntity.SessionToken.IsValid() {
			c.String(http.StatusUnauthorized, "expired credential provided")
			return
		}

		c.Set("user", userEntity)
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
