package product_routes

import (
	"encoding/json"
	"net/http"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	product_utils "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/utils"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/product", "GET", readProductRoute)
	println("Route registered: readProduct")
}

func readProductRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var prod product.ProductStock

	// Build the filter
	var nameRaw, _ = utils.GetQueryParam(r, "name", false, "")
	var idRaw, _ = utils.GetQueryParam(r, "id", false, "")

	filter := product_utils.NewQueryFilter()
	if nameRaw != "" {
		filter.SetName(nameRaw)
	}
	if idRaw != "" {
		filter.SetID(idRaw)
	}
	var query bson.M
	if query, err = filter.Build(); err != nil {
		http.Error(w, "ID received invalid: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Perform the query
	if prod, err = app.GetInstance().Services.Product.Read(query); err != nil {
		http.Error(w, "Some error occurred while querying: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prod); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
