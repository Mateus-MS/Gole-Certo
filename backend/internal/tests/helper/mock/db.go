package tests_mock

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(t *testing.T) *mongo.Database {
	t.Helper()

	// Create an unique name to this DB
	dbName := fmt.Sprintf("testdb_%s", t.Name())

	client, err := mongo.Connect(t.Context(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database(dbName)

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = client.Database(dbName).Drop(ctx)
		_ = client.Disconnect(ctx)
	})

	return db
}
