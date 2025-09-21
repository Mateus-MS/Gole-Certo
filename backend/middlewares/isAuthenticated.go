package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	user_repository "alves.com/backend/modules/users/repo"
	user_service "alves.com/backend/modules/users/service"
	"github.com/gin-gonic/gin"
)

var (
	ErrMissingAuthHeader = errors.New("missing authorization header")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
	ErrExpiredToken      = errors.New("expired token provided")
)

// NOTE: this function only validates the received token by checking it's cached or armazened duration
// This should really be improved since this way it's just a "weak authentication" and opens a lot of security flaws
func AuthMiddleware(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, ErrMissingAuthHeader) || errors.Is(err, ErrInvalidAuthHeader) {
				status = http.StatusBadRequest
			}
			c.String(status, err.Error())
			return
		}

		expiresAt, err := getTokenExpiry(c, token, userService)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, ErrExpiredToken) || errors.Is(err, user_repository.ErrUserInexistent) {
				status = http.StatusUnauthorized
			}
			c.String(status, err.Error())
			return
		}

		if time.Now().After(expiresAt) {
			c.String(http.StatusUnauthorized, ErrExpiredToken.Error())
			return
		}

		c.Next()
	}
}

func getTokenExpiry(ctx context.Context, token string, service user_service.IService) (time.Time, error) {
	// Try Redis first
	expiresAt, err := getExpiryFromCache(ctx, token, service)
	if err == nil {
		return expiresAt, err
	}

	// If not in Redis, try on Mongo
	user, err := service.Repo().ReadBySessionToken(ctx, token)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserInexistent) {
			return time.Time{}, ErrExpiredToken
		}
		return time.Time{}, err
	}

	// Since it's not in cache, save on Redis
	_ = service.Cache().Set(ctx, token, user.SessionToken.ExpiresAt, 30*time.Minute)

	return user.SessionToken.ExpiresAt, nil
}

func getExpiryFromCache(ctx context.Context, token string, service user_service.IService) (time.Time, error) {
	expiresStr, err := service.Cache().Get(ctx, token)
	if err != nil {
		return time.Time{}, err
	}

	expiresAt, err := time.Parse(time.RFC3339, expiresStr)
	if err != nil {
		return time.Time{}, err
	}

	return expiresAt, nil
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
