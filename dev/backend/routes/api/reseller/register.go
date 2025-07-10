package api_reseller

import (
	"encoding/json"
	"net/http"

	cpf "github.com/Mateus-MS/Gole-Certo/dev/backend/models/CPF"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/models/email"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

type test struct {
	Name  string `json:"Name"`
	CPF   string `json:"CPF"`
	Email string `json:"Email"`
}

func init() {
	app.GetInstance().Router.RegisterRoutes("/api/reseller", "POST", registerSellerRoute)
	println("Route registered: registerSeller")
}

func registerSellerRoute(w http.ResponseWriter, r *http.Request) {
	var err error
	var test test

	if err = json.NewDecoder(r.Body).Decode(&test); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var clientCpf cpf.CPF
	if clientCpf, err = cpf.New(test.CPF); err != nil {
		http.Error(w, "Invalid CPF: "+err.Error(), http.StatusBadRequest)
		return
	}

	var clientEmail email.Email
	if clientEmail, err = email.New(test.Email); err != nil {
		http.Error(w, "Invalid Email: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(clientEmail.Get() + " " + clientCpf.Get()))
}

func registerIndividual() {

}

func registerCompany() {

}
