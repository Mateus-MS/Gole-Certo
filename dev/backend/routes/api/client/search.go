package api_client

import (
	"encoding/json"
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/client", "GET", searchClientRoute)
	println("Route registered: searchClient")
}

func searchClientRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var client client.Client
	var identifierRaw string

	// TODO: sanitize
	if identifierRaw, err = utils.GetQueryParam(r, "identifier", true, ""); err != nil {
		http.Error(w, "Request must include the client identifier in parameters", http.StatusBadRequest)
		return
	}

	// client.Search returns a NOT concrete type, it returns the Interface client.Client
	if client, err = app.GetInstance().Repositories.Client.Search(identifierRaw); err != nil {
		http.Error(w, "Some error occurred while querying: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
