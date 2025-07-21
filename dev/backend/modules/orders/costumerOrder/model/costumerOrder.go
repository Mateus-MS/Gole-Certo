package costumerOrder

import (
	"errors"
	"slices"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyProductList = errors.New("empty product list")
	ErrInvalidState     = errors.New("invalid state")
)

type CostumerOrder struct {
	ID       primitive.ObjectID `json:"ID"        bson:"_id"`
	Products []product.Product  `json:"Products"  bson:"products"`
	State    string             `json:"State"     bson:"state"`

	// TODO: see if worth it stores the other user data here like adresses
	UserID string `json:"UserID" bson:"userID"`
}

func New(prods []product.Product, state string) (CostumerOrder, error) {
	if len(prods) <= 0 {
		return CostumerOrder{}, ErrEmptyProductList
	}

	if !slices.Contains([]string{"batching", "processing", "delivering", "received"}, state) {
		return CostumerOrder{}, ErrInvalidState
	}

	return CostumerOrder{
		Products: prods,
		State:    state,
	}, nil
}
