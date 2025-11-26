package costumerOrder_service

import (
	"context"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
)

func (s *service) Register(ctx context.Context, ord costumerOrder.CostumerOrder) (ordID string, err error) {
	// Check if the received user exists on DB
	if !s.userService.Repo().HasUser(ctx, ord.UserID) {
		return "", err
	}

	// check if the received products list really exists on DB
	for _, prod := range ord.Products {
		stock, err := s.stockService.Repo().ReadByID(ctx, prod.GetProductID())
		if err != nil {
			return "", err
		}

		// perhaps the desirable behavior is not this
		// maybe be flatten the order
		// or don't do nothing allowing this to continue normally
		if stock.Stock < prod.GetAmmount() {
			return "", ErrInsufficientStock
		}

		// Remove from stock what was ordered
		if err := s.stockService.DeductFromStock(ctx, stock, prod.GetAmmount()); err != nil {
			return "", err
		}
	}

	// Create the order
	err = s.Repo().Create(ctx, ord)
	if err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}
