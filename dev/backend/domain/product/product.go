package product

type Product struct {
	ProductID string  `json:"ProductID" bson:"productID"`
	Price     float32 `json:"Price"     bson:"price"`
	Quantity  int     `json:"Quantity"  bson:"quantity"`
}
