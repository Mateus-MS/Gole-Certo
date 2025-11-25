package stock_error

import "errors"

var (
	ErrStockInexistent    = errors.New("this product does not exists on DB")
	ErrStockAlreadyExists = errors.New("this product already exists on DB")
)
