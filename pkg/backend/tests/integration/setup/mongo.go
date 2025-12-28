package integration_setup

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Name     string
}

func NewTestDB(tName string) (*TestDB, error) {
	ctx := context.Background()

	// Connect
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}

	// Generate a random name
	dbName := fmt.Sprintf(
		"test_%s_%d",
		tName,
		time.Now().UnixNano(),
	)

	db := client.Database(dbName)

	return &TestDB{
		Client:   client,
		Database: db,
		Name:     dbName,
	}, nil
}

func (tDB *TestDB) Teardown() error {
	ctx := context.Background()

	// Drop entire database
	if err := tDB.Database.Drop(ctx); err != nil {
		return err
	}

	// Close connection
	return tDB.Client.Disconnect(ctx)
}
