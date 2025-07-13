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
	var typeRaw string
	var cli client.Client

	if typeRaw, err = utils.GetQueryParam(r, "type", true, ""); err != nil {
		http.Error(w, "Request must include the client type in parameters", http.StatusBadRequest)
		return
	}

	if typeRaw == "individual" {
		var ind client.Individual

		if err = json.NewDecoder(r.Body).Decode(&ind); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ind.Type = typeRaw
		cli = &ind
	}

	if typeRaw == "company" {
		var comp client.Company

		if err = json.NewDecoder(r.Body).Decode(&comp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		comp.Type = typeRaw
		cli = &comp
	}

	// TODO: show in the error message the parameter that is missing
	if !cli.IsValid() {
		http.Error(w, "Request must be a complete client", http.StatusBadRequest)
		return
	}

	if err = app.GetInstance().Repositories.Client.Save(cli); err != nil {
		http.Error(w, "Failed to save the client into DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
