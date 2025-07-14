package app

import (
	"net/http"
	"sync"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository/mock"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository/persistence"
	"go.mongodb.org/mongo-driver/mongo"
)

var app_instance *Application
var app_once sync.Once

type Application struct {
	DB           *mongo.Client
	Router       *Router
	Repositories *repository.Repositories
}

func GetInstance() *Application {
	app_once.Do(func() {
		app_instance = newApplication()
	})
	return app_instance
}

func newApplication() *Application {
	// Create the router
	router := CreateRouter()

	// Serve static files from the "frontend" directory
	router.Mux.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("dev/frontend"))))

	db := StartDBConnection()

	repositories := repository.Repositories{
		Client:  &persistence.ClientRepository{DB: db},
		Product: &mock.ProductRepository{DB: db},
		Order:   &persistence.OrderRepository{DB: db},
	}

	// Return the application instance
	return &Application{
		DB:           db,
		Router:       &router,
		Repositories: &repositories,
	}
}
