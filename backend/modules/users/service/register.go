package user_service

import (
	"context"

	user_repository "alves.com/backend/modules/users/repo"
)

func (s *service) Register(ctx context.Context, username, password string) error {
	// check if already exists an user with this name
	userEntity, err := s.Repo().ReadByName(ctx, username)
	if err == nil {
		return user_repository.ErrDuplicatedUser
	}

	userEntity.Name = username

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	userEntity.Password = hashedPassword

	// Save into DB
	err = s.repository.Create(ctx, &userEntity)
	if err != nil {
		return err
	}

	return nil
}
