package user_service

import (
	"context"
	"time"

	user_model "alves.com/backend/modules/users/model"
	user_repository "alves.com/backend/modules/users/repo"
)

func (s *service) Login(ctx context.Context, username, password string) error {
	// Search for the user on DB
	userEntity, err := s.repository.ReadByName(ctx, username)
	if err != nil {
		return user_repository.ErrUserInexistent
	}

	// Check if the finded user password, match with the received one
	if !CheckPassword(userEntity.Password, password) {
		return ErrInvalidCredentials
	}

	// Generate a session token
	sessionToken, err := user_model.NewToken(20, 30*time.Minute)
	if err != nil {
		return err
	}

	// Stores the generated tokens into the queried DB entity
	userEntity.SessionToken = sessionToken

	// Update the DB state of user with the generated token
	err = s.repository.UpdateByName(ctx, userEntity)
	if err != nil {
		return err
	}

	// Add the token to the cache
	err = s.cache.Set(ctx, sessionToken.Token, sessionToken.ExpiresAt, 30*time.Minute)
	if err != nil {
		return err
	}

	return nil
}
