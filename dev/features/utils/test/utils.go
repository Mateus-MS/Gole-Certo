package testutils

import (
	"context"
	"testing"
)

func SetupTest(t *testing.T) *Application {
	app := createTestApp()
	teardown(t, app)

	return app
}

func teardown(t *testing.T, app *Application) {
	t.Cleanup(func() {
		app.DB.Database("MOCK").Drop(context.TODO())
	})
}
