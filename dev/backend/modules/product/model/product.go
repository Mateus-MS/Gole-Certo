package product

import (
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrInvalidBrand    = errors.New("invalid brand")
	ErrInvalidID       = errors.New("invalid id")

	ErrDuplicated        = errors.New("product already registered")
	ErrProductInexistent = errors.New("product does not exists on db")
)

type Product struct {
	ProductID primitive.ObjectID   `json:"ProductID,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"Name"                bson:"name"`
	Brand     string               `json:"Brand"               bson:"brand"`
	Price     primitive.Decimal128 `json:"Price"               bson:"price"`
	Quantity  int64                `json:"Quantity"            bson:"quantity"`
}

// Constructor
func New(name, brand, priceStr string, quantity int64) (Product, error) {
	// Converts the received price into correct format
	var price primitive.Decimal128
	var err error
	if price, err = primitive.ParseDecimal128(priceStr); err != nil {
		return Product{}, ErrInvalidPrice
	}

	prod := Product{
		ProductID: primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:      name,
		Brand:     brand,
		Price:     price,
		Quantity:  quantity,
	}

	if err := prod.IsValid(); err != nil {
		return prod, err
	}

	return prod, nil
}

// Validator
func (p *Product) IsValid() error {
	if p.Name == "" {
		return ErrInvalidName
	}
	if p.Brand == "" {
		return ErrInvalidBrand
	}

	// Check if the price is lower than 1
	priceFloat, err := strconv.ParseFloat(p.Price.String(), 64)
	if err != nil {
		return ErrInvalidPrice
	}
	if priceFloat <= 0 {
		return ErrInvalidPrice
	}

	// Check the quantity
	if p.Quantity < 0 {
		return ErrInvalidQuantity
	}

	return nil
}

func (p Product) Equals(other Product) bool {
	return p.ProductID == other.ProductID
}

// Returns a merged list with
// added up quantities of any repeated product
// and append any unique
func MergeLists(p1, p2 []Product) []Product {
	merged := append([]Product{}, p1...)

	for _, newProd := range p2 {
		found := false
		for i, existing := range merged {
			if newProd.Equals(existing) {
				merged[i].Quantity += newProd.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, newProd)
		}
	}

	return merged
}
