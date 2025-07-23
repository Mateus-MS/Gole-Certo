package app

import (
	"net/http"
	"sync"

	duffbeer_service "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer"
	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer/mock"
	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	stock_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/service"
	user_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var app_instance *Application
var app_once sync.Once

type Application struct {
	DB       *mongo.Client
	Router   *Router
	Services *Services
}

type Services struct {
	User  contracts.User_Service
	Stock contracts.Stock_Service

	SupplierOrder contracts.SupplierOrder_Service
	CostumerOrder contracts.CostumerOrder_Service

	DuffBeer duffbeer_service.Service
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

func createServices(client *mongo.Client) *Services {
	user := user_service.New(client.Database("goleCertoDB").Collection("users"))
	prod := stock_service.New(client.Database("goleCertoDB").Collection("stock"))

	duffbeer := duffbeerService_mock.New()

	return &Services{
		User:     user,
		Stock:    prod,
		DuffBeer: duffbeer,
	}
}
