package integration_fixtures

import (
	"context"

	"alves.com/backend/internal/security"
	user_model "alves.com/backend/modules/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTestUser(ctx context.Context, db *mongo.Database, username, password string) primitive.ObjectID {
	// Hash the pass
	passHashed, err := security.HashPassword(password)
	if err != nil {
		panic(err)
	}

	// Create the user
	user := user_model.NewUser(username, passHashed)

	// Store into DB
	_, err = db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}

	return user.ID
}

func InsertTestAdmUser(ctx context.Context, db *mongo.Database) primitive.ObjectID {
	// Hash the pass
	passHashed, err := security.HashPassword("adm")
	if err != nil {
		panic(err)
	}

	// Create the user
	user := user_model.NewUser("adm", passHashed)
	user.IsAdmin = true

	// Store into DB
	_, err = db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}

	return user.ID
}
