package productservice

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"

type Service interface {
	Search(string) (product.Product, error)
}
