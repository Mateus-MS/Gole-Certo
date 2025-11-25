package stock_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StockEntity struct {
	ID   primitive.ObjectID `json:"ID" bson:"_id"`
	Name string             `json:"name" binding:"required"`
}
