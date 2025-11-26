package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Create some sort of custom structs like did with router to be simplier to deal with DB

func StartDBConnection() (mongoClient *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	if mongoClient, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		log.Fatal("Mongo connection error: " + err.Error())
	}

	return mongoClient
}
