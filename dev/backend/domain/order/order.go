package order

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"

// Order state can be:
// - `batching` when the orders are batching up, till all products in orders in DB with this same states sums up to 1000
// - `proceed` when successfully ordered in Duff beer systems

type Order struct {
	OrderID        string            `json:"OrderID"  bson:"_id"`
	UserIdentifier string            `json:"UserID"   bson:"userID"`
	Product        []product.Product `json:"Products" bson:"products"`
	State          string            `json:"State"    bson:"state"`
}

func New(userID, orderID, state string, products []product.Product) Order {
	return Order{
		UserIdentifier: userID,
		Product:        products,
		OrderID:        orderID,
		State:          state,
	}
}
