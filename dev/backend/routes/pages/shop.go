package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	page_shop "github.com/Mateus-MS/Gole-Certo/dev/frontend/pages/shop"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/shop", "GET", ShopPage)
	println("Page registered: ShopPage")
}

func ShopPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// If no page index provided
	if query.Get("page") == "" {
		// Sent with 1
		query.Set("page", "1")
	}

	// Rebuild raw query
	rawQuery := query.Encode()

	// Get the brands from the request
	filters, err := utils.GetProductFilters(r, true, false)
	if err != nil {
		http.Error(w, "Something went wrong while querying the brands in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	brands, err := app.GetInstance().Services.Stock.Repo().GetBrands(r.Context(), filters)
	if err != nil {
		http.Error(w, "Something went wrong while querying the brands in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Simple re-pass the request params to the HTMX inside Shop Page
	// The endpoint of prodPage htmx component will validate the parameters
	page_shop.ShopPage(rawQuery, brands).Render(r.Context(), w)
}
