package user_repository_mongo

import (
	"context"
	"errors"

	generic_repository "alves.com/backend/modules/common/repo"
	user_model "alves.com/backend/modules/users/model"
	user_repository "alves.com/backend/modules/users/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByName(ctx context.Context, name string) (user_model.UserEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"name": name})

	if err != nil {
		if errors.Is(err, generic_repository.ErrItemInexistent) {
			return user_model.UserEntity{}, user_repository.ErrUserInexistent
		} else {
			return user_model.UserEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	user, ok := userGeneric.(*user_model.UserEntity)
	if !ok {
		return user_model.UserEntity{}, user_repository.ErrCannotConvert
	}

	return *user, nil
}
