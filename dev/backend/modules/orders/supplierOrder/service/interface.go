package supplierOrder_service

import (
	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
)

type Order = supplierOrder.SupplierOrder

type Service interface {
	// Through the repository layer It will:
	// - Check all products received are valids
	// - Check if there is some batch... well, batching
	// - Will update the existing batch with the sended products
	// - If no batch, creates a new one
	Register(Order) (string, error)

	ReadByOrderID(any) (Order, error)
	ReadOneByState(string) (Order, error)
	ReadManyByState(string, int) ([]Order, error)

	create(Order) (string, error)
	UpdateByID(Order) error
}
