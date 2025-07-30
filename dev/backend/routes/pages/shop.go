package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	page_shop "github.com/Mateus-MS/Gole-Certo/dev/frontend/pages/shop"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/shop", "GET", ShopPage)
	println("Page registered: ShopPage")
}

func ShopPage(w http.ResponseWriter, r *http.Request) {
	prods, err := app.GetInstance().Services.Stock.Repo().ReadManyFilteredAfterID(r.Context(), bson.M{}, "", 12)
	if err != nil {
		http.Error(w, "Error while querying the procuts to show onto screen", http.StatusInternalServerError)
		return
	}

	page_shop.ShopPage(prods).Render(r.Context(), w)
}
