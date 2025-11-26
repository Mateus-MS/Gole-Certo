package stock_service

import (
	"context"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
)

// NOTE: This function normally will be called ONLY after user and prod validation
// so i will not validate that again here, if needed refactor later
func (s *service) DeductFromStock(ctx context.Context, prod Stock, quantityToRemove int64) error {
	// Deduct from stock
	prod.Stock -= quantityToRemove

	// Update it on DB
	if err := s.Repo().UpdateByID(ctx, prod); err != nil {
		return err
	}

	// Check for the Min threshold
	reStockAmount := prod.CalculateRestockAmount()
	if reStockAmount == 0 {
		// No need to re-stock
		return nil
	}

	// Need to re-stock
	supOrder, err := s.createNewOrder(prod, reStockAmount)
	if err != nil {
		return err
	}

	// Register the new order
	_, err = s.supplierOrder.Register(ctx, supOrder)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) createNewOrder(prod Stock, reStockAmount int64) (supplierOrder.SupplierOrder, error) {
	supProd := prod.GetInSupplierFormat()
	supProd.Quantity = reStockAmount

	// Create the order OBJ
	supOrder, err := supplierOrder.New(
		[]*supplierOrder.SupplierProduct{
			supProd,
		},
	)
	if err != nil {
		return supplierOrder.SupplierOrder{}, err
	}

	return supOrder, nil
}
