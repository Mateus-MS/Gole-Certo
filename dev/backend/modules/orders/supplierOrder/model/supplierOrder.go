package supplierOrder

import (
	"errors"
	"slices"
	"time"

	product "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/product/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrEmptyProductList = errors.New("empty product list")
	ErrInvalidState     = errors.New("invalid state")
)

type SupplierOrder struct {
	ID       primitive.ObjectID `json:"ID"        bson:"_id"`
	Products []product.Product  `json:"Products"  bson:"products"`
	State    string             `json:"State"     bson:"state"`
}

func New(prods []product.Product, state string) (SupplierOrder, error) {
	if len(prods) <= 0 {
		return SupplierOrder{}, ErrEmptyProductList
	}

	if !slices.Contains([]string{"batching", "processing", "delivering", "delivered"}, state) {
		return SupplierOrder{}, ErrInvalidState
	}

	return SupplierOrder{
		ID:       primitive.NewObjectIDFromTimestamp(time.Now()),
		Products: prods,
		State:    state,
	}, nil
}
