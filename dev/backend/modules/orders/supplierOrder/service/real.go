package supplierOrder_service

import (
	"errors"

	contracts "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/common"
	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
	supplierOrder_repository "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/repository"
	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/stock/model"
	"github.com/Mateus-MS/Gole-Certo/dev/features/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrOrderStateMustBeBatching = errors.New("supplier order status must be batching")
)

type service struct {
	repository supplierOrder_repository.Repository

	// Dependencies
	stockService contracts.Stock_Service
}

func (s *service) SetStockService(stockService contracts.Stock_Service) {
	s.stockService = stockService
}

func New(coll *mongo.Collection) service {
	return service{
		repository: *supplierOrder_repository.New(coll),
	}
}

func (s *service) Register(ord supplierOrder.SupplierOrder) (_ string, err error) {
	// Check if all products received, are valids
	for _, prod := range ord.Products {
		if !s.stockService.ValidateProductByID(prod.GetProductID()) {
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
		return s.handleBatching(ord)
	}

	// Otherwise just create it
	if err := s.repository.Create(ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}

func (s *service) ReadByOrderID(ordID_any any) (ord supplierOrder.SupplierOrder, err error) {
	var ordID_obj primitive.ObjectID
	if ordID_obj, err = utils.ParseObjectID(ordID_any); err != nil {
		return ord, err
	}

	if ord, err = s.repository.Read(bson.M{"_id": ordID_obj}); err != nil {
		return ord, err
	}

	return ord, nil
}

func (s *service) ReadOneByState(state string) (supplierOrder.SupplierOrder, error) {
	filter := bson.M{"state": state}

	orders, err := s.repository.Read(filter)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) ReadManyByState(state string, limit int) ([]supplierOrder.SupplierOrder, error) {
	filter := bson.M{"state": state}

	orders, err := s.repository.ReadMany(filter, limit)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *service) UpdateByID(updateState supplierOrder.SupplierOrder) (err error) {
	// Get the actual state of the order
	var realState supplierOrder.SupplierOrder
	if realState, err = s.ReadByOrderID(updateState.ID); err != nil {
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
	if err = s.repository.Update(updateState, bson.M{"_id": updateState.ID}); err != nil {
		return err
	}

	return nil
}

func (s *service) create(ord supplierOrder.SupplierOrder) (_ string, err error) {
	// Save into DB
	if err := s.repository.Create(ord); err != nil {
		return "", err
	}

	return ord.ID.Hex(), nil
}

func (s *service) handleBatching(updateState supplierOrder.SupplierOrder) (_ string, err error) {
	// Check if there is any batch available
	var realState supplierOrder.SupplierOrder
	if realState, err = s.ReadOneByState("batching"); err != nil {
		// There is no Batch ready on DB
		if err == supplierOrder_repository.ErrOrderNotFound {
			//  Create a new one
			var neworderid string
			if neworderid, err = s.create(updateState); err != nil {
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
	err = s.UpdateByID(updateState)
	if err != nil {
		return "", err
	}

	return updateState.ID.Hex(), nil
}
