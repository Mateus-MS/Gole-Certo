package supplierOrder

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierProduct struct {
	Name      string             `json:"Name"      bson:"name"`
	Quantity  int64              `json:"Quantity"  bson:"quantity"`
	ProductID primitive.ObjectID `json:"-"         bson:"-"`
}

func (p *SupplierProduct) GetProductID() string {
	return p.ProductID.Hex()
}

func (p *SupplierProduct) SetAmmount(q int64) {
	p.Quantity = q
}
func (p *SupplierProduct) GetAmmount() int64 {
	return p.Quantity
}
