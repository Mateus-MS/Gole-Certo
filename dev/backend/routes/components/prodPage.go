package components

import (
	"net/http"
	"strconv"

	stock_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/repository/mongo"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"github.com/Mateus-MS/Gole-Certo/dev/frontend/components"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/components/prodPage", "GET", prodPageRoute)
	println("Component registered: ProdPage")
}

func prodPageRoute(w http.ResponseWriter, r *http.Request) {
	pageIndexRaw, err := utils.GetQueryParam(r, "page", true, "")
	if err != nil {
		http.Error(w, "Must pass the page index", http.StatusBadRequest)
		return
	}
	pageIndex, err := strconv.Atoi(pageIndexRaw)
	if err != nil {
		http.Error(w, "The page index must be a number", http.StatusBadRequest)
		return
	}

	prods, err := app.GetInstance().Services.Stock.Repo().ReadManyPaged(r.Context(), bson.M{}, int64(pageIndex), stock_repository.ItemsPerPage)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	totalPages, err := app.GetInstance().Services.Stock.Repo().TotalPages(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Render the templ component to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := components.ProductPageComponent(prods, int64(pageIndex), totalPages).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
