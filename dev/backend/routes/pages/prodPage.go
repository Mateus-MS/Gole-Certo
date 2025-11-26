package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	page_prodpage "github.com/Mateus-MS/Gole-Certo/dev/frontend/pages/prodPage"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/productpage", "GET", prodPageRoute)
	println("Page registered: ProductPage")
}

func prodPageRoute(w http.ResponseWriter, r *http.Request) {

	productID, err := utils.GetQueryParam(r, "id", true, "")
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		return
	}

	productData, err := app.GetInstance().Services.Stock.Repo().ReadByID(r.Context(), productID)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		return
	}

	page_prodpage.ProdPage(productData).Render(r.Context(), w)
}
