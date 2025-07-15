package productservice

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"
	"github.com/Mateus-MS/Gole-Certo/dev/features/app"
)

func Search(identifier string) (prod product.Product, err error) {
	return app.GetInstance().Repositories.Product.Search(identifier)
}
