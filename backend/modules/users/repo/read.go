package user_repository

import (
	"context"
	"errors"

	generic_repository "alves.com/modules/common/repo"
	user_model "alves.com/modules/users/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByName(ctx context.Context, name string) (user_model.UserEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"name": name})

	if err != nil {
		if errors.Is(err, generic_repository.ErrItemInexistent) {
			return user_model.UserEntity{}, ErrUserInexistent
		} else {
			return user_model.UserEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	user, ok := userGeneric.(*user_model.UserEntity)
	if !ok {
		return user_model.UserEntity{}, ErrCannotConvert
	}

	return *user, nil
}

func (repo *Repository) ReadBySessionToken(ctx context.Context, token string) (user_model.UserEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"sessionToken.token": token})

	if err != nil {
		if errors.Is(err, generic_repository.ErrItemInexistent) {
			return user_model.UserEntity{}, ErrUserInexistent
		} else {
			return user_model.UserEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	user, ok := userGeneric.(*user_model.UserEntity)
	if !ok {
		return user_model.UserEntity{}, ErrCannotConvert
	}

	return *user, nil
}
