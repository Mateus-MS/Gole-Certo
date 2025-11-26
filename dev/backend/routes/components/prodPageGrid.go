package components

import (
	"context"
	"net/http"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
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
	// TODO: VALIDATE/SANITAZE THE REQUEST PARAMS

	// Get the page index from the url
	pageIndex, err := utils.GetQueryParam(r, "page", true, 1)
	if err != nil {
		http.Error(w, "Must pass the page index", http.StatusBadRequest)
		return
	}

	filter := product.StockFilter{}
	if err = filter.ReadBrands(r); err != nil {
		println(err.Error())
		http.Error(w, "Something went wrong while building the brands filter: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err = filter.ReadPriceRange(r); err != nil {
		println(err.Error())
		http.Error(w, "Something went wrong while building the priceRange filter: "+err.Error(), http.StatusBadRequest)
		return
	}

	filterBson := filter.Build()

	// Get the product repository
	prodRepository := app.GetInstance().Services.Stock.Repo()

	ascending, err := utils.GetQueryParam(r, "price-order", false, "ascending")
	if err != nil {
		http.Error(w, "Must pass the page index", http.StatusBadRequest)
		return
	}

	// Get the data to render the page
	var prods []stock_repository.Product
	if prods, err = getProducts(r.Context(), prodRepository, int64(pageIndex), filterBson, ascending == "ascending"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var totalPages int64
	if totalPages, err = getTotalPages(r.Context(), prodRepository, filterBson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the templ component to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := components.ProductPageGridComponent(filterBson, prods, int64(pageIndex), totalPages).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func getProducts(context context.Context, repo *stock_repository.Repository, pageIndex int64, filter bson.M, ascending bool) ([]stock_repository.Product, error) {
	prods, err := repo.ReadManyPaged(context, filter, pageIndex, stock_repository.ItemsPerPage, ascending)
	if err != nil {
		return nil, err
	}

	return prods, nil
}

func getTotalPages(context context.Context, repo *stock_repository.Repository, filter bson.M) (int64, error) {
	totalPages, err := repo.TotalPages(context, filter)
	if err != nil {
		return 0, err
	}

	return totalPages, nil
}
