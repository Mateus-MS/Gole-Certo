package user_service

import (
	"context"

	user_repository "alves.com/backend/modules/users/repo"
)

func (s *service) Login(ctx context.Context, username, password string) error {
	// Search for the user on DB
	userEntity, err := s.repository.ReadByName(ctx, username)
	if err != nil {
		return user_repository.ErrUserInexistent
	}

	// Check if the finded user password, match with the received one
	if !passwordMatch(userEntity.Password, password) {
		return ErrInvalidCredentials
	}

	// Generate a session token
	sessionToken, err := GenerateToken(20)
	if err != nil {
		return err
	}

	// Set the generated token in the queried user state
	userEntity.SessionToken = sessionToken

	// Update the DB state of user with the generated token
	s.repository.UpdateByName(ctx, userEntity)

	return nil
}

func passwordMatch(entityPass, requestPass string) bool {
	return requestPass == entityPass
}
