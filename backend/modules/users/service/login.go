package user_service

import (
	"context"
	"time"

	user_repository "alves.com/modules/users/repo"
)

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	// Search for the user on DB
	userEntity, err := s.repository.ReadByName(ctx, username)
	if err != nil {
		return "", user_repository.ErrUserInexistent
	}

	// Check if the finded user password, match with the received one
	if !CheckPassword(userEntity.Password, password) {
		return "", ErrInvalidCredentials
	}

	// Generate a session token
	sessionToken, err := GenerateRandomToken(20)

	// Add the token to the cache
	err = s.cache.Set(ctx, sessionToken, userEntity.ID.Hex(), 30*time.Minute)
	if err != nil {
		println("CACHE ERROR", err.Error())
		return "", err
	}

	return sessionToken, nil
}
