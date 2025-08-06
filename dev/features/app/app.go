package app

import (
	"net/http"
	"sync"

	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer/mock"
	costumerOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/service"
	supplierOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/service"
	product_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/service"
	user_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/service"

	duffbeer_service "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer"
	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
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
	stock := product_service.New(client.Database("goleCertoDB").Collection("stock"))
	supplierOrder := supplierOrder_service.New(client.Database("goleCertoDB").Collection("supplier_orders"))
	costumerOrder := costumerOrder_service.New(client.Database("goleCertoDB").Collection("costumer_orders"))

	// Add the dependecies
	supplierOrder.SetStockService(stock)

	costumerOrder.SetStockService(stock)
	costumerOrder.SetUserService(user)

	stock.SetSupplierOrderService(&supplierOrder)

	duffbeer := duffbeerService_mock.New()

	return &Services{
		User:  user,
		Stock: stock,

		SupplierOrder: &supplierOrder,
		CostumerOrder: &costumerOrder,

		DuffBeer: duffbeer,
	}
}
