package order

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/product"

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
