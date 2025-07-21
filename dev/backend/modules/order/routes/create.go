package api_order

import (
	"encoding/json"
	"net/http"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/order", "POST", registerOrder)
	println("Route registered: registerOrder")
}

func registerOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		ordID string

		// Anounymous structs
		request struct {
			Products []product.Product `json:"Products"`
			UserID   string            `json:"UserID"`
		}
		response struct {
			Products []product.Product `json:"Products"`
			OrderID  string            `json:"OrderID"`
		}
	)

	// 1- Build the request json
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error while decoding the received JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 2 - Save the order in DB
	if ordID, err = app.GetInstance().Services.Order.Create(request.UserID, request.Products); err != nil {
		http.Error(w, "Error while registering order in DB: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 3- Build the response
	response.OrderID = ordID
	response.Products = request.Products

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
