package components

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"github.com/Mateus-MS/Gole-Certo/dev/frontend/components"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/components/filtersSuggestions", "GET", filtersSuggestionsRoute)
	println("Component registered: Filters Suggestions")
}

func filtersSuggestionsRoute(w http.ResponseWriter, r *http.Request) {
	search, err := utils.GetQueryParam(r, "search", false, "")
	if err != nil {
		http.Error(w, "Something went wrong while querying the brands in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the brands from the request
	filters, err := utils.GetProductFilters(r, true, false)
	if err != nil {
		http.Error(w, "Something went wrong while querying the brands in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	filters["limit"] = 5
	filters["search"] = search

	brands, err := app.GetInstance().Services.Stock.Repo().GetBrands(r.Context(), filters)
	if err != nil {
		http.Error(w, "Something went wrong while querying the brands in DB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the templ component to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := components.FiltersSuggestionsComponent(brands).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
