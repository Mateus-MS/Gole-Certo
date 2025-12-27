package order_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderEntity struct {
	ID       primitive.ObjectID         `json:"ID" bson:"_id"`
	UserID   primitive.ObjectID         `json:"UserID" bson:"userId"`
	Products map[primitive.ObjectID]int `json:"Products"  bson:"products"`
}

// Used to unmarshal the json expected from the HTTP request
type OrderRequest struct {
	UserID   string         `json:"UserID"`
	Products map[string]int `json:"products"`
}

func New(userID primitive.ObjectID, products map[primitive.ObjectID]int) *OrderEntity {
	return &OrderEntity{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		UserID:   userID,
		Products: products,
	}
}
