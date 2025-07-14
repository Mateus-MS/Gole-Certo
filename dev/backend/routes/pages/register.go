package routes_pages

import (
	"net/http"

	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
	page_register "github.com/Mateus-MS/Gole-Certo/dev/frontend/pages/register"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/register", "GET", RegisterPage)
	println("Page registered: RegisterPage")
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	page_register.RegisterPage().Render(r.Context(), w)
}
