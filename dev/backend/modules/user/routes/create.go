package api_user

import (
	"encoding/json"
	"net/http"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/user", "POST", registerUserRoute)
	println("Route registered: registerUser")
}

func registerUserRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var typeRaw string
	var usr user.User

	// Get the type received from the param
	if typeRaw, err = utils.GetQueryParam(r, "type", true, ""); err != nil {
		http.Error(w, "Request must include the user type in parameters", http.StatusBadRequest)
		return
	}

	// Unmarshal the JSON into the struct accordantly with the type
	if typeRaw == "individual" {
		var ind user.Individual

		if err = json.NewDecoder(r.Body).Decode(&ind); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ind.Type = typeRaw
		usr = &ind
	}

	if typeRaw == "company" {
		var comp user.Company

		if err = json.NewDecoder(r.Body).Decode(&comp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		comp.Type = typeRaw
		usr = &comp
	}

	// Save in DB
	if err = app.GetInstance().Services.User.Register(r.Context(), usr); err != nil {
		http.Error(w, "Failed to save the user into DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
