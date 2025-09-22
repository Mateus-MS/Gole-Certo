package user_repository

import (
	"context"

	user_model "alves.com/modules/users/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) UpdateByName(ctx context.Context, user user_model.UserEntity) error {
	set := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"password": user.Password,
		},
	}
	return repo.Update(ctx, bson.M{"name": user.Name}, set)
}
