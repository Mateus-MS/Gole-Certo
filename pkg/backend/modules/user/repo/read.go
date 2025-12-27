package user_repository

import (
	"context"
	"errors"

	user_error "alves.com/backend/modules/user/errors"
	user_model "alves.com/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Repository) Read(ctx context.Context, filter bson.M) (user_model.UserEntity, error) {
	var user user_model.UserEntity

	err := repo.Collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, user_error.ErrUserInexistent
		}
	}

	return user, nil
}

func (repo *Repository) ReadByName(ctx context.Context, name string) (*user_model.UserEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"name": name})

	if err != nil {
		if errors.Is(err, user_error.ErrUserInexistent) {
			return &user_model.UserEntity{}, user_error.ErrUserInexistent
		} else {
			return &user_model.UserEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	return &userGeneric, nil
}

func (repo *Repository) ReadByID(ctx context.Context, id primitive.ObjectID) (*user_model.UserEntity, error) {
	userGeneric, err := repo.Read(ctx, bson.M{"_id": id})

	if err != nil {
		if errors.Is(err, user_error.ErrUserInexistent) {
			return &user_model.UserEntity{}, user_error.ErrUserInexistent
		} else {
			return &user_model.UserEntity{}, errors.Join(errors.New("something went wrong"), err)
		}
	}

	return &userGeneric, nil
}
