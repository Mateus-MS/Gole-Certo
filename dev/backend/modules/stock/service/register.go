package stock_service

import (
	"context"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
)

func (s *service) Register(ctx context.Context, prod Stock) (err error) {
	if err = prod.IsValid(); err != nil {
		return err
	}

	// Search for any product with the same `name` into DB
	// NOTE: should not have two products with same name
	if _, err = s.repository.ReadByName(ctx, prod.Name); err == nil {
		return product.ErrDuplicated
	}

	return s.repository.Create(ctx, prod)
}
