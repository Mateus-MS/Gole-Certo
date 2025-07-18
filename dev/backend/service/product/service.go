package productservice

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"

type Service interface {
	Create(product.Product) error
	Read(QueryFilter) (product.Product, error)
}
