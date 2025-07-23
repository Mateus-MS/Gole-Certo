package testutils

import (
	"net/http"

	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer/mock"
	costumerOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/service"
	supplierOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/service"
	product_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/service"
	user_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/service"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	DB       *mongo.Client
	Router   *http.ServeMux
	Services *app.Services
}

func createTestApp() *Application {
	mux := http.NewServeMux()
	db := app.StartDBConnection()

	user := user_service.New(db.Database("MOCK").Collection("users"))
	stock := product_service.New(db.Database("MOCK").Collection("stock"))
	supplierOrder := supplierOrder_service.New(db.Database("MOCK").Collection("supplier_orders"))
	costumerOrder := costumerOrder_service.New(db.Database("MOCK").Collection("costumer_orders"))

	// Add the dependecies
	supplierOrder.SetStockService(stock)

	costumerOrder.SetStockService(stock)
	costumerOrder.SetUserService(user)

	duffbeer := duffbeerService_mock.New()

	services := &app.Services{
		User:  user,
		Stock: stock,

		SupplierOrder: &supplierOrder,
		CostumerOrder: &costumerOrder,

		DuffBeer: duffbeer,
	}

	// Return the application instance
	return &Application{
		DB:       db,
		Router:   mux,
		Services: services,
	}

}
