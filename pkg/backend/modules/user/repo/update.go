package user_repository

import (
	"context"

	user_error "alves.com/backend/modules/user/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) Update(ctx context.Context, filter bson.M, updateSet bson.M) error {
	result, err := repo.Collection.UpdateOne(ctx, filter, updateSet)

	if result.MatchedCount == 0 {
		return user_error.ErrUserInexistent
	}

	if err != nil {
		return err
	}

	return nil
}
