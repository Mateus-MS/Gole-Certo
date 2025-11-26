package api_user

import (
	"encoding/json"
	"net/http"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/user", "GET", searchUserRoute)
	println("Route registered: searchUser")
}

func searchUserRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var usr user.User
	var identifierRaw string

	// TODO: sanitize
	if identifierRaw, err = utils.GetQueryParam(r, "identifier", true, ""); err != nil {
		http.Error(w, "Request must include the user identifier in parameters", http.StatusBadRequest)
		return
	}

	// user.Search returns a NOT concrete type, it returns the Interface user.User
	if usr, err = app.GetInstance().Services.User.Repo().ReadByID(r.Context(), identifierRaw); err != nil {
		http.Error(w, "Some error occurred while querying: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
