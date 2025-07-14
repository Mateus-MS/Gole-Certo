package order

import (
	"encoding/json"
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/order"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/google/uuid"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/order", "POST", registerOrder)
	println("Route registered: registerOrder")
}

func registerOrder(w http.ResponseWriter, r *http.Request) {
	reps := app.GetInstance().Repositories
	var (
		err error
		ord order.Order

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

	// 2- Check if the received user identifier, identifies any user
	// TODO: Sanitize this input
	if _, err = reps.Client.Search(request.UserID); err != nil {
		http.Error(w, "User ID in request does not exists in DB. "+err.Error(), http.StatusNotFound)
		return
	}

	// 3- Check if the received product list match existing products
	// NOTE: currently, it's not checking, it's using a mock, always returning true :P
	for _, product := range request.Products {
		if _, err := app.GetInstance().Repositories.Product.Search(product.ProductID); err != nil {
			http.Error(w, "Product does not exist: "+product.ProductID, http.StatusNotFound)
			return
		}
	}

	// 4- Save the order in the DB
	ord = order.New(
		request.UserID,      // UserID
		uuid.New().String(), // OrderID
		"processing",        // State
		request.Products,    // Products
	)

	if err = reps.Order.Save(ord); err != nil {
		http.Error(w, "Error while saving order in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 5- Build the response
	response.OrderID = ord.OrderID
	response.Products = request.Products

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
