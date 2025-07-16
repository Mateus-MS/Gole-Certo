package productservice_mock

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) Search(identifier string) (prod product.Product, err error) {
	return prod, err
}
