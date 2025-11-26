package product_routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/product", "POST", registerProduct)
	println("Route registered: registerProduct")
}

// TODO: Only allow from ADMs
func registerProduct(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		prod product.ProductStock
	)

	// 1- Build the request json
	if err = json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, "Error while decoding the received JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Temp
	prod.ProductID = primitive.NewObjectIDFromTimestamp(time.Now())

	// 2 - Save the product in DB
	if err = app.GetInstance().Services.Stock.Register(r.Context(), prod); err != nil {
		if errors.Is(err, product.ErrDuplicated) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Error while registering order in DB: "+err.Error(), http.StatusBadRequest)
		return
	}
}
