package supplierOrder

import (
	"errors"
	"slices"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyProductList       = errors.New("empty product list")
	ErrInvalidState           = errors.New("invalid state")
	ErrInvalidProductQuantity = errors.New("invalid quantity")
)

type SupplierOrder struct {
	ID            primitive.ObjectID `json:"ID"             bson:"_id"`
	Products      []*SupplierProduct `json:"Products"       bson:"products"`
	State         string             `json:"State"          bson:"state"`
	TotalQuantity int64              `json:"TotalQuantity"  bson:"totalQuantity"`
}

func New(prods []*SupplierProduct, state string) (SupplierOrder, error) {
	if len(prods) <= 0 {
		return SupplierOrder{}, ErrEmptyProductList
	}

	// TODO: think if is better to just use "batching" as default state instead of receiving it
	if !slices.Contains([]string{"batching", "processing", "delivering", "delivered"}, state) {
		return SupplierOrder{}, ErrInvalidState
	}

	return SupplierOrder{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Products: prods,
		State:    state,
	}, nil
}
