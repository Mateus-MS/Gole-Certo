package user_service

import (
	"context"
	"time"

	user_repository "alves.com/modules/users/repo"
	"alves.com/pkg/security"
)

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	// Search for the user on DB
	userEntity, err := s.ReadByName(ctx, username)
	if err != nil {
		return "", user_repository.ErrUserInexistent
	}

	// Check if the finded user password, match with the received one
	if !security.CheckPassword(userEntity.Password, password) {
		return "", ErrInvalidCredentials
	}

	// Generate a session token
	sessionToken, err := security.GenerateRandomToken(20)

	// Create the user cache entity
	userCache := userEntity.GetCache()

	// Add the token to the cache
	err = s.SaveInCache(ctx, sessionToken, *userCache, 30*time.Minute)
	if err != nil {
		println("CACHE ERROR", err.Error())
		return "", err
	}

	return sessionToken, nil
}
