package user_service

import (
	"context"

	user_model "alves.com/modules/users/model"
	user_repository "alves.com/modules/users/repo"
)

func (s *service) Register(ctx context.Context, username, password string) error {
	// check if already exists an user with this name
	_, err := s.Repo().ReadByName(ctx, username)
	if err == nil {
		return user_repository.ErrDuplicatedUser
	}

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Save into DB
	err = s.repository.Create(
		ctx,
		user_model.NewUser(username, hashedPassword),
	)
	if err != nil {
		return err
	}

	return nil
}
