package stock_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StockEntity struct {
	ID       primitive.ObjectID `json:"ID" bson:"_id"`
	Name     string             `json:"name" binding:"required"`
	Quantity int                `json:"quantity"`
}

func New(name string, quantity int) *StockEntity {
	return &StockEntity{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:     name,
		Quantity: quantity,
	}
}
