package api_reseller

import (
	"encoding/json"
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/reseller", "POST", registerSellerRoute)
	println("Route registered: registerSeller")
}

func registerSellerRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var clientRaw client.Individual

	if err = json.NewDecoder(r.Body).Decode(&clientRaw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Re-send the json just for test :P
	if err := json.NewEncoder(w).Encode(clientRaw); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
