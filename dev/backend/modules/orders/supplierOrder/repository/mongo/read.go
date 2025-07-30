package supplierOrder_repository

import (
	"context"

	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *Repository) ReadByOrderID(ctx context.Context, ordID_any any) (ord supplierOrder.SupplierOrder, err error) {
	var ordID_obj primitive.ObjectID
	if ordID_obj, err = utils.ParseObjectID(ordID_any); err != nil {
		return ord, err
	}

	if ord, err = repo.read(ctx, bson.M{"_id": ordID_obj}); err != nil {
		return ord, err
	}

	return ord, nil
}

func (repo *Repository) ReadOneByState(ctx context.Context, state string) (supplierOrder.SupplierOrder, error) {
	orders, err := repo.read(ctx, bson.M{"state": state})
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (repo *Repository) ReadManyByState(ctx context.Context, state string, limit int) ([]supplierOrder.SupplierOrder, error) {
	orders, err := repo.readMany(ctx, bson.M{"state": state}, limit)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
