package costumerOrder

import "go.mongodb.org/mongo-driver/bson/primitive"

type CostumerProduct struct {
	ProductID primitive.ObjectID `json:"ProductID" bson:"productID"`
	Quantity  int64              `json:"Quantity"  bson:"quantity"`
}

func (p *CostumerProduct) GetProductID() string {
	return p.ProductID.Hex()
}

func (p *CostumerProduct) SetAmmount(q int64) {
	p.Quantity = q
}
func (p *CostumerProduct) GetAmmount() int64 {
	return p.Quantity
}
