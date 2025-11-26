package supplierOrder_service

import (
	"context"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository/mongo"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
)

func (s *service) Register(ctx context.Context, ord supplierOrder.SupplierOrder) (_ string, err error) {
	// Check if all products received, are valids
	for _, prod := range ord.Products {
		if !s.stockService.Repo().HasProduct(ctx, prod.GetProductID()) {
			return "", product.ErrProductInexistent
		}
	}

	// Count how many products are
	var prodCount int64
	for _, prod := range ord.Products {
		prodCount += prod.Quantity
	}
	ord.TotalQuantity = prodCount

	// Handle the batching case
	if ord.State == "batching" {
		return s.handleBatching(ctx, ord)
	}

	// Otherwise just create it
	if _, err := s.repository.Create(ctx, ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}

func (s *service) handleBatching(ctx context.Context, updateState supplierOrder.SupplierOrder) (_ string, err error) {
	// Check if there is any batch available
	var realState supplierOrder.SupplierOrder
	if realState, err = s.Repo().ReadOneByState(ctx, "batching"); err != nil {
		// There is no Batch ready on DB
		if err == supplierOrder_repository.ErrOrderNotFound {
			//  Create a new one
			var neworderid string
			if neworderid, err = s.Repo().Create(ctx, updateState); err != nil {
				// Error during the batch creation
				return "", err
			}
			return neworderid, nil
		}
		// Other error has occurred
		return "", err
	}

	// Update the existing one
	updateState.ID = realState.ID
	err = s.Repo().UpdateByID(ctx, updateState)
	if err != nil {
		return "", err
	}

	return updateState.ID.Hex(), nil
}
