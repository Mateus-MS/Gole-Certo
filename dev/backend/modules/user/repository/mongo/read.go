package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) ReadByID(ctx context.Context, identifier string) (c User, err error) {
	// Build the filter
	filter := bson.M{"_id": identifier}

	// Query
	raw, err := repo.read(ctx, filter)
	if err != nil {
		return c, err
	}

	// Decode
	return decodeUser(raw)

}
