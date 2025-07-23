package product

import (
	"errors"
	"strconv"
	"time"

	costumerOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/costumerOrder/model"
	supplierOrder "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/orders/supplierOrder/model"
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

type ProductStock struct {
	ProductID    primitive.ObjectID   `json:"ProductID,omitempty" bson:"_id,omitempty"`
	Name         string               `json:"Name"                bson:"name"`
	Brand        string               `json:"Brand"               bson:"brand"`
	Price        primitive.Decimal128 `json:"Price"               bson:"price"`
	Stock        int64                `json:"Stock"               bson:"stock"`
	MinThreshold int8                 `json:"MinThreshold"        bson:"minthreshold"`
	MaxStock     int64                `json:"MaxStock"            bson:"maxstock"`
}

// Constructor
func New(name, brand, priceStr string, stock int64) (ProductStock, error) {
	// Converts the received price into correct format
	var price primitive.Decimal128
	var err error
	if price, err = primitive.ParseDecimal128(priceStr); err != nil {
		return ProductStock{}, ErrInvalidPrice
	}

	prod := ProductStock{
		ProductID:    primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:         name,
		Brand:        brand,
		Price:        price,
		Stock:        stock,
		MinThreshold: 15, // 15%
		MaxStock:     200,
	}

	if err := prod.IsValid(); err != nil {
		return prod, err
	}

	return prod, nil
}

// Validator
func (p *ProductStock) IsValid() error {
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
	if p.Stock < 0 {
		return ErrInvalidQuantity
	}

	// Check the quantity
	// Since doesn't make sense the MIN threshold be 90%
	if p.MinThreshold <= 0 && p.MinThreshold >= 90 {
		return ErrInvalidQuantity
	}

	return nil
}

func (p *ProductStock) GetProductID() string {
	return p.ProductID.Hex()
}

func (p *ProductStock) SetAmmount(q int64) {
	p.Stock = q
}
func (p *ProductStock) GetAmmount() int64 {
	return p.Stock
}

// Utils
func (p *ProductStock) CalculateRestockAmount() int64 {
	threshold := int64(p.MinThreshold) * p.MaxStock / 100
	if p.Stock < threshold {
		return p.MaxStock - p.Stock
	}
	return 0
}

// Converters

func (p *ProductStock) GetInSupplierFormat() *supplierOrder.SupplierProduct {
	return &supplierOrder.SupplierProduct{
		Name:      p.Name,
		ProductID: p.ProductID,
	}
}

func (p *ProductStock) GetInCostumerFormat() *costumerOrder.CostumerProduct {
	return &costumerOrder.CostumerProduct{
		ProductID: p.ProductID,
	}
}
