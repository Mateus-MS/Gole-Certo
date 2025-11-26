package costumerOrder

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyProductList = errors.New("empty product list")
	ErrInvalidState     = errors.New("invalid state")
)

type CostumerOrder struct {
	ID       primitive.ObjectID `json:"ID"        bson:"_id"`
	Products []CostumerProduct  `json:"Products"  bson:"products"`
	State    string             `json:"State"     bson:"state"`

	// TODO: see if worth it stores the other user data here like adresses
	UserID string `json:"UserID" bson:"userID"`
}

func New(prods []CostumerProduct) (CostumerOrder, error) {
	if len(prods) <= 0 {
		return CostumerOrder{}, ErrEmptyProductList
	}

	return CostumerOrder{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Products: prods,
		State:    "processing",
	}, nil
}
