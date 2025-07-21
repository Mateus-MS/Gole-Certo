package testutils

import (
	"net/http"

	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/external/duffbeer/mock"
	supplierOrder_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/service"
	product_service "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/service"
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
	prod := product_service.New(db.Database("MOCK").Collection("products"))
	ordr := supplierOrder_service.New(
		db.Database("MOCK").Collection("supplier_orders"),
		prod,
	)
	duffbeer := duffbeerService_mock.New()

	services := &app.Services{
		User:    user,
		Product: prod,

		SupplierOrder: &ordr,

		DuffBeer: duffbeer,
	}

	// Return the application instance
	return &Application{
		DB:       db,
		Router:   mux,
		Services: services,
	}

}
