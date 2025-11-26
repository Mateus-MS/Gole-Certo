package supplierOrder_repository

import (
	"context"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) UpdateByID(ctx context.Context, updateState supplierOrder.SupplierOrder) (err error) {
	// Get the actual state of the order
	var realState supplierOrder.SupplierOrder
	if realState, err = repo.ReadByOrderID(ctx, updateState.ID); err != nil {
		return err
	}

	// Merge the DB state with the updated one
	updateState.Products = product.MergeLists(updateState.Products, realState.Products)

	// Count how many products are
	var prodCount int64
	for _, prod := range updateState.Products {
		prodCount += prod.Quantity
	}
	updateState.TotalQuantity = prodCount

	// Save it
	if err = repo.update(ctx, updateState, bson.M{"_id": updateState.ID}); err != nil {
		return err
	}

	return nil
}
