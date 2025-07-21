package order

import (
	"errors"
	"slices"
	"time"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidOrderID   = errors.New("invalid order id")
	ErrInvalidState     = errors.New("invalid state")
	ErrEmptyProductList = errors.New("empty product list")
)

type Order struct {
	OrderID primitive.ObjectID `json:"OrderID,omitempty" bson:"_id,omitempty"`
	Product []product.Product  `json:"Products"          bson:"products"`
	State   string             `json:"State"             bson:"state"`
}

func New(state string, products []product.Product) (ord Order, err error) {
	if len(products) <= 0 {
		return ord, ErrEmptyProductList
	}

	if !slices.Contains([]string{"batching", "ordered"}, state) {
		return ord, ErrInvalidState
	}

	return Order{
		OrderID: primitive.NewObjectIDFromTimestamp(time.Now()),
		Product: products,
		State:   state,
	}, nil
}
