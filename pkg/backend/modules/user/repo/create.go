package user_repository

import (
	"context"

	user_error "alves.com/backend/modules/user/errors"
	user_model "alves.com/backend/modules/user/model"
)

func (repo *Repository) Create(ctx context.Context, user user_model.UserEntity) error {
	_, err := repo.Collection.InsertOne(ctx, user)

	if err != nil {
		return user_error.ErrUserAlreadyExists
	}

	return nil
}
