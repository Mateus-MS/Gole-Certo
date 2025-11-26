package user_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Works almost identically as `ReadById` but instead of returning the user or some error, it only returns a bool
func (repo *Repository) HasUser(ctx context.Context, identifier string) bool {
	filter := bson.M{"_id": identifier}

	err := repo.collection.FindOne(ctx, filter).Err()

	return err == nil || !errors.Is(err, mongo.ErrNoDocuments)
}
