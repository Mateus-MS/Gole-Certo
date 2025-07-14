package product

type Product struct {
	ProductID string `json:"ProductID" bson:"productID"`
	Quantity  int    `json:"Quantity"  bson:"quantity"`
}
