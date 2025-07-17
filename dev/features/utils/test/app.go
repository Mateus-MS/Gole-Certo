package testutils

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/repository"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/service"
	duffbeerService_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/service/external/duffbeer/mock"
	orderservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/order"
	productservice_mock "github.com/Mateus-MS/Gole-Certo/dev/backend/service/product/mock"
	userservice "github.com/Mateus-MS/Gole-Certo/dev/backend/service/user"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	DB       *mongo.Client
	Router   *http.ServeMux
	Services *service.Services
}

func createTestApp() *Application {
	mux := http.NewServeMux()
	db := app.StartDBConnection()

	user := userservice.New(repository.UserRepository{Collection: db.Database("goleCertoDB_MOCK").Collection("users")})
	prod := productservice_mock.New()
	ordr := orderservice.New(
		repository.OrderRepository{Collection: db.Database("goleCertoDB_MOCK").Collection("orders")},
		user,
		prod,
	)
	duffbeer := duffbeerService_mock.New()

	services := &service.Services{
		User:     user,
		Product:  prod,
		Order:    ordr,
		DuffBeer: duffbeer,
	}

	// Return the application instance
	return &Application{
		DB:       db,
		Router:   mux,
		Services: services,
	}

}
