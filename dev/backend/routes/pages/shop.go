package routes_pages

import (
	"math"
	"net/http"
	"strconv"

	stock_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/repository/mongo"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	page_shop "github.com/Mateus-MS/Gole-Certo/dev/frontend/pages/shop"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/shop", "GET", ShopPage)
	println("Page registered: ShopPage")
}

func ShopPage(w http.ResponseWriter, r *http.Request) {

	// Get the page index received in the request
	pageIndexRaw, err := utils.GetQueryParam(r, "page", false, "1")
	if err != nil {
		http.Error(w, "Something went wrong :P", http.StatusInternalServerError)
		return
	}

	pageIndex, err := strconv.Atoi(pageIndexRaw)
	if err != nil {
		http.Error(w, "The received page index must be a valid integer", http.StatusBadRequest)
		return
	}

	if pageIndex < 1 {
		pageIndex = 1
	}

	// Get the total quantity of items in DB
	// TODO: it can be optimized
	total, err := app.GetInstance().Services.Stock.Repo().TotalItems(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, "Something went wrong while counting the quantity of items in DB", http.StatusInternalServerError)
		return
	}

	// Calculate the quantity of pages
	totalPages := int(math.Ceil(float64(total) / float64(stock_repository.ItemsPerPage)))

	if pageIndex > totalPages {
		pageIndex = totalPages
	}

	page_shop.ShopPage(strconv.Itoa(pageIndex)).Render(r.Context(), w)
}
