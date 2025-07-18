package app

import (
	"net/http"
	"sync"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/service"
	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/service/external/duffbeer/mock"
	orderservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/order"
	productservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var app_instance *Application
var app_once sync.Once

type Application struct {
	DB       *mongo.Client
	Router   *Router
	Services *service.Services
}

func GetInstance() *Application {
	app_once.Do(func() {
		app_instance = newApplication()
	})
	return app_instance
}

func newApplication() *Application {
	// Create the router
	router := createRouter()

	// Serve static files from the "frontend" directory
	router.Mux.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("dev/frontend"))))

	db := StartDBConnection()

	// Return the application instance
	return &Application{
		DB:       db,
		Router:   &router,
		Services: createServices(db),
	}
}

func createServices(client *mongo.Client) *service.Services {
	user := userservice.New(repository.UserRepository{Collection: client.Database("goleCertoDB").Collection("users")})
	prod := productservice.New(repository.ProductRepository{Collection: client.Database("goleCertoDB").Collection("products")})
	ordr := orderservice.New(
		repository.OrderRepository{Collection: client.Database("goleCertoDB").Collection("orders")},
		user,
		prod,
	)
	duffbeer := duffbeerService_mock.New()

	return &service.Services{
		User:     user,
		Product:  prod,
		Order:    ordr,
		DuffBeer: duffbeer,
	}
}
