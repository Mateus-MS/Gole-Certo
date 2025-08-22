package components

import (
	"context"
	"net/http"
	"strconv"
	"strings"

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

	// Build the filter to query the data
	var filter bson.M
	if filter, err = getFilters(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the product repository
	prodRepository := app.GetInstance().Services.Stock.Repo()

	// Get the data to render the page
	var prods []stock_repository.Product
	if prods, err = getProducts(r.Context(), prodRepository, int64(pageIndex), filter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var totalPages int64
	if totalPages, err = getTotalPages(r.Context(), prodRepository, filter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the templ component to the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := components.ProductPageComponent(filter, prods, int64(pageIndex), totalPages).Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func getFilters(r *http.Request) (filter bson.M, err error) {
	filter = bson.M{}

	// Get the "brand" filter from the URL
	var brandRaw string
	if brandRaw, err = utils.GetQueryParam(r, "brands", false, ""); err != nil {
		return filter, err
	}
	if brandRaw != "" {
		// Assume multiple brands are separated by commas
		brands := strings.Split(brandRaw, ":")
		filter["brand"] = bson.M{
			"$in": brands,
		}
	}

	// Get the "price" filter from the URL
	var priceRaw string
	if priceRaw, err = utils.GetQueryParam(r, "price", false, ""); err != nil {
		return filter, err
	}
	if priceRaw != "" {
		priceMinMax := strings.Split(priceRaw, "-")
		if len(priceMinMax) == 2 {
			min, _ := strconv.ParseFloat(priceMinMax[0], 64)
			max, _ := strconv.ParseFloat(priceMinMax[1], 64)
			filter["price"] = bson.M{
				"$gte": min,
				"$lte": max,
			}
		}
	}

	return filter, nil
}

func getProducts(context context.Context, repo *stock_repository.Repository, pageIndex int64, filter bson.M) ([]stock_repository.Product, error) {
	prods, err := repo.ReadManyPaged(context, filter, pageIndex, stock_repository.ItemsPerPage)
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
