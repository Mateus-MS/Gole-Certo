package integration_fixtures

import (
	"context"

	stock_model "alves.com/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTestStock(ctx context.Context, db *mongo.Database, name string, quantity int) primitive.ObjectID {
	// Create the user
	stock := stock_model.New(name, quantity)

	// Store into DB
	_, err := db.Collection("stock").InsertOne(ctx, stock)
	if err != nil {
		panic(err)
	}

	return stock.ID
}
