package api_client

import (
	"encoding/json"
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/client", "POST", registerSellerRoute)
	println("Route registered: registerSeller")
}

func registerSellerRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var client client.Individual
	var typeRaw string

	if err = json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if typeRaw, err = utils.GetQueryParam(r, "type", true, ""); err != nil {
		http.Error(w, "Request must include the client type in parameters", http.StatusBadRequest)
		return
	}
	client.Type = typeRaw

	if err = app.GetInstance().Repositories.Client.Save(&client); err != nil {
		http.Error(w, "Failed to save the client into DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
